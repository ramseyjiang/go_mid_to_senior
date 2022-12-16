package httpreqtest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Constants
const sumMethodName = "Sum"
const multiplyMethodName = "Multiply"

type AppHandlerFake struct {
	// method name -> call -> params
	Calls map[string][][]interface{}
}

func (a *AppHandlerFake) Sum(x int, y int) (r Result) {
	b := a.Calls[sumMethodName]
	c := []interface{}{x, y}
	a.Calls[sumMethodName] = append(b, c)

	r.Value = 7 // assert result, it should be the same with input params
	return
}

func (a *AppHandlerFake) Multiply(x int, y int) (r Result) {
	b := a.Calls[multiplyMethodName]
	c := []interface{}{x, y}

	a.Calls[multiplyMethodName] = append(b, c)

	r.Value = 10 // assert result, it should be the same with input params
	return
}

// Create our test table
var testTable = []struct {
	// The name of the test
	name string
	// The HTTP method to use in our call
	method string
	// The URL path that is being requested
	path string
	// The expected response status code
	statusCode int
	// The expected response body, as string
	body string
	// The request body to sent with the request
	requestBody map[string]interface{}
	// The name of the AppHandlerFake method that we want to spy on
	handlerMethodName string
	// The parameters we expect the 'handlerMethodName' on the AppHandlerFake to be called with
	handlerToBeCalledWith []interface{}
	// The headers that are being set for the request
	requestHeaders map[string]string
	// The response headers we want to test on
	headers map[string]string
}{
	{
		name:                  `GET endpoint to get a sum`,
		method:                http.MethodGet,
		path:                  `/sum?x=5&y=2`,
		statusCode:            http.StatusOK,
		requestBody:           nil,
		body:                  `{"value":7}`,
		handlerMethodName:     sumMethodName,
		handlerToBeCalledWith: []interface{}{5, 2},
		headers:               map[string]string{`Content-Type`: `application/json`},
	},
	{
		name:       `POST endpoint to multiply, wrong header`,
		method:     http.MethodPost,
		path:       `/multiply`,
		statusCode: http.StatusInternalServerError,
		requestBody: map[string]interface{}{
			"x": 2,
			"y": 3,
		},
		body:           `unprocessable Entity`,
		requestHeaders: map[string]string{`Content-Type`: `application/text`},
		headers:        map[string]string{`Content-Type`: `text/plain; charset=utf-8`},
	},
	{
		name:       `POST endpoint to multiply`,
		method:     http.MethodPost,
		path:       `/multiply`,
		statusCode: http.StatusOK,
		requestBody: map[string]interface{}{
			"x": 4,
			"y": 5,
		},
		body:                  `{"value":10}`,
		handlerMethodName:     multiplyMethodName,
		handlerToBeCalledWith: []interface{}{4, 5},
		headers:               map[string]string{`Content-Type`: `application/json`},
	},
}

func TestApp(t *testing.T) {
	appHandler := &AppHandlerFake{}
	app := CreateApp(appHandler)

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			appHandler.Calls = map[string][][]interface{}{}

			// Create and send request
			reqBody, _ := json.Marshal(tc.requestBody)
			request := httptest.NewRequest(tc.method, tc.path, bytes.NewReader(reqBody))
			request.Header.Add(`Content-Type`, `application/json`)

			// Request Headers
			for k, v := range tc.requestHeaders {
				request.Header.Add(k, v)
			}

			response, _ := app.Test(request)
			if response.StatusCode != tc.statusCode {
				t.Errorf("StatusCode was incorrect, got: %d, want: %d.", response.StatusCode, tc.statusCode)
			}

			// Headers
			for k, want := range tc.headers {
				if response.Header.Get(k) != want {
					t.Errorf("Response header '%s' was incorrect, got: '%s', want: '%s'", k, response.Header.Get(k), want)
				}
			}

			// Response Body
			body, _ := io.ReadAll(response.Body)
			actual := string(body)
			if actual != tc.body {
				t.Errorf("Body was incorrect, got: %v, want: %v", actual, tc.body)
			}

			// Check if handler was called correctly
			if tc.handlerMethodName != "" {
				if !cmp.Equal(appHandler.Calls[tc.handlerMethodName][0], tc.handlerToBeCalledWith) {
					t.Errorf("Handler method '%s' wasn't called with the correct parameters. Got: '%v', want: '%v'", tc.handlerMethodName, appHandler.Calls[tc.handlerMethodName][0], tc.handlerToBeCalledWith)
				}
			}
		})
	}
}
