package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

type closeNotifyingRecorder struct {
	*httptest.ResponseRecorder
	closeNotifyChan chan bool
}

func newCloseNotifyingRecorder() *closeNotifyingRecorder {
	return &closeNotifyingRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}

func (cnr *closeNotifyingRecorder) CloseNotify() <-chan bool {
	return cnr.closeNotifyChan
}

func TestCreateRouter(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		method   string
		path     string
		expected string
	}{
		{
			name:     "GetUserList",
			method:   "GET",
			path:     "/user/list",
			expected: "Service 1",
		},
		{
			name:     "AddUser",
			method:   "POST",
			path:     "/user/create",
			expected: "Service 2",
		},
		{
			name:     "LogAdd",
			method:   "POST",
			path:     "/log/add",
			expected: "Service 3",
		},
	}

	// Activate the httpmock library
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the servers
	httpmock.RegisterResponder("GET", "http://localhost:8080/user/list",
		httpmock.NewStringResponder(200, "Service 1"))
	httpmock.RegisterResponder("POST", "http://localhost:8080/user/create",
		httpmock.NewStringResponder(200, "Service 2"))
	httpmock.RegisterResponder("POST", "http://localhost:8081/log/add",
		httpmock.NewStringResponder(200, "Service 3"))

	// Create the router
	router := createRouter()

	// Run test cases
	for _, tc := range testCases {
		tc := tc // Capture range variable

		t.Run(tc.name, func(t *testing.T) {
			// Create a request
			req, _ := http.NewRequest(tc.method, tc.path, nil)
			resp := newCloseNotifyingRecorder()

			// Serve the request
			router.ServeHTTP(resp, req)

			// Check the status code
			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Equal(t, tc.expected, resp.Body.String())
		})
	}
}
