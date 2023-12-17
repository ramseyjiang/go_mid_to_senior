package loginrate

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestRateLimitMiddleware(t *testing.T) {
	rl := NewRateLimiter()

	tests := []struct {
		name           string
		username       string
		requestCount   int
		expectedStatus int
	}{
		{"UnderLimit", "user1", 4, http.StatusOK},
		{"AtLimit", "user2", 5, http.StatusOK},
		{"OverLimit", "user3", 6, http.StatusTooManyRequests},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := RateLimitMiddleware(rl, http.HandlerFunc(LoginHandler))

			for i := 1; i <= tt.requestCount; i++ {
				// Create form data
				formData := url.Values{}
				formData.Set("username", tt.username)

				// Create a request with form data
				req := httptest.NewRequest("POST", "/login", strings.NewReader(formData.Encode()))
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

				rr := httptest.NewRecorder()

				handler.ServeHTTP(rr, req)

				// Check the response status of the last request
				if i > maxLoginAttempts {
					if status := rr.Code; status != tt.expectedStatus {
						t.Errorf("handler returned wrong status code: got %v want %v",
							status, tt.expectedStatus)
					}
				}
			}

			// Clear the entry for the next test
			rl.mutex.Lock()
			delete(rl.loginAttempts, tt.username)
			rl.mutex.Unlock()
		})
	}
}
