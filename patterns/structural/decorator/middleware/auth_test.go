package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// creates an HTTP request and a response recorder, and then uses the ApplyMiddleware function to wrap the HelloHandler with the LoggingMiddleware and AuthenticationMiddleware functions. The resulting HTTP handler is then called with the request and response recorder to simulate a real HTTP request. Finally, the test checks that the response status code is http.StatusForbidden, which is the expected result when the Authorization header is not set to "secret_token".
	t.Run("Middleware Auth - no authorization", func(t *testing.T) {
		rr := httptest.NewRecorder()
		handler := ApplyMiddleware(HelloHandler, LoggingMiddleware, AuthenticationMiddleware)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusForbidden {
			t.Errorf("unexpected status code: got %v want %v", status, http.StatusForbidden)
		}
	})

	// creates an HTTP request and a response recorder, sets the Authorization header to "secret_token", and then uses the ApplyMiddleware function to wrap the HelloHandler with the LoggingMiddleware and AuthenticationMiddleware functions. The resulting HTTP handler is then called with the request and response recorder to simulate a real HTTP request. Finally, the test checks that the response status code is http.StatusOK and the response body is equal to "Hello, World!\n", which is the expected result when the Authorization header is set to "secret_token".
	t.Run("Middleware Auth - authorization", func(t *testing.T) {
		req.Header.Set("Authorization", "secret_token")
		rr := httptest.NewRecorder()
		handler := ApplyMiddleware(HelloHandler, LoggingMiddleware, AuthenticationMiddleware)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("unexpected status code: got %v want %v", status, http.StatusOK)
		}

		expected := "Hello, World!\n"
		if rr.Body.String() != expected {
			t.Errorf("unexpected response body: got %v want %v", rr.Body.String(), expected)
		}
	})
}
