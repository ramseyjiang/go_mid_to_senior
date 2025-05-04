package reverseproxy

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
	// Activate the httpmock library
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the servers
	httpmock.RegisterResponder("GET", "http://localhost:8081/test",
		httpmock.NewStringResponder(200, "Service 1"))
	httpmock.RegisterResponder("GET", "http://localhost:8082/test",
		httpmock.NewStringResponder(200, "Service 2"))

	// Create the router
	router := createRouter()

	// Create a request to the first route
	req, _ := http.NewRequest("GET", "/service1/test", nil)
	resp := newCloseNotifyingRecorder()

	// Serve the request
	router.ServeHTTP(resp, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Service 1", resp.Body.String())

	// Create a request to the second route
	req, _ = http.NewRequest("GET", "/service2/test", nil)
	resp = newCloseNotifyingRecorder()

	// Serve the request
	router.ServeHTTP(resp, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Service 2", resp.Body.String())
}
