package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name       string
		route      string
		method     string
		handler    func(w http.ResponseWriter, r *http.Request)
		wantStatus int
		wantBody   string
	}{
		{
			name:       "TestHomeHandler",
			route:      "/",
			method:     "GET",
			handler:    homeHandler,
			wantStatus: http.StatusOK,
			wantBody:   "Welcome to the Home Page!",
		},
		{
			name:       "TestContactHandlerGET",
			route:      "/connect",
			method:     "GET",
			handler:    contactHandler,
			wantStatus: http.StatusOK,
			wantBody:   "GET request to the Contact Page!",
		},
		{
			name:       "TestContactHandlerPOST",
			route:      "/connect",
			method:     "POST",
			handler:    contactHandler,
			wantStatus: http.StatusOK,
			wantBody:   "POST request to the Contact Page!",
		},
		{
			name:       "TestUserHandler",
			route:      "/user?id=123",
			method:     "GET",
			handler:    userHandler,
			wantStatus: http.StatusOK,
			wantBody:   "User ID: 123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(tt.handler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("%v returned wrong status code: got %v want %v", tt.name, status, tt.wantStatus)
			}

			if rr.Body.String() != tt.wantBody {
				t.Errorf("%v returned unexpected body: got %v want %v", tt.name, rr.Body.String(), tt.wantBody)
			}
		})
	}
}
