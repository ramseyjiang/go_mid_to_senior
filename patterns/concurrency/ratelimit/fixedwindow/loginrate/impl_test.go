package loginrate

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateLimitMiddleware(t *testing.T) {
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
			handler := RateLimitMiddleware(http.HandlerFunc(LoginHandler))

			for i := 0; i < tt.requestCount; i++ {
				req := httptest.NewRequest("POST", "/login", nil)
				req.Form.Set("username", tt.username)
				rr := httptest.NewRecorder()

				handler.ServeHTTP(rr, req)

				if i == tt.requestCount-1 {
					if status := rr.Code; status != tt.expectedStatus {
						t.Errorf("handler returned wrong status code: got %v want %v",
							status, tt.expectedStatus)
					}
				}
			}

			// Reset the count for the next test
			mutex.Lock()
			delete(loginAttempts, tt.username)
			mutex.Unlock()
		})
	}
}
