package server

import (
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestWebServer(t *testing.T) {
	webServer := newWebServer()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	httpCode, body := webServer.handleRequest(appStatusURL, "GET")
	assert.Equal(t, http.StatusOK, httpCode)
	assert.Equal(t, "Ok", body)

	httpCode, body = webServer.handleRequest(appStatusURL, "GET")
	assert.Equal(t, http.StatusOK, httpCode)
	assert.Equal(t, "Ok", body)

	httpCode, body = webServer.handleRequest(appStatusURL, "GET")
	assert.Equal(t, http.StatusForbidden, httpCode) // because the maxAllowedRequest in the server is 2.
	assert.Equal(t, "Not Allowed", body)

	httpCode, body = webServer.handleRequest(createUserURL, "POST")
	assert.Equal(t, http.StatusCreated, httpCode)
	assert.Equal(t, "User Created", body)

	httpCode, body = webServer.handleRequest(createUserURL, "GET")
	assert.Equal(t, http.StatusNotFound, httpCode)
	assert.Equal(t, "Not Ok", body)
}
