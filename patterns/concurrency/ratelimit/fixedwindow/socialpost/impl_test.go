package socialpost

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimitMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		requestCount   int
		expectedStatus int
		resetAfter     time.Duration
	}{
		{"UnderLimit", "user1", 20, http.StatusOK, 0},
		{"AtLimit", "user2", 30, http.StatusOK, 0},
		{"OverLimit", "user3", 31, http.StatusTooManyRequests, 0},
		{"ResetAfterOneMinute", "user4", 31, http.StatusOK, 1 * time.Minute},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			req, err := http.NewRequest("GET", "/post", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("X-User-ID", tt.userID)

			for i := 0; i < tt.requestCount; i++ {
				rr := httptest.NewRecorder()
				handler := RateLimitMiddleware(http.HandlerFunc(PostHandler))

				// Act
				handler.ServeHTTP(rr, req)

				// Assert
				if i == tt.requestCount-1 {
					if status := rr.Code; status != tt.expectedStatus {
						t.Errorf("handler returned wrong status code: got %v want %v",
							status, tt.expectedStatus)
					}
				}
			}

			// Reset the count for the next test
			if tt.resetAfter > 0 {
				time.Sleep(tt.resetAfter)
				mutex.Lock()
				delete(userActivities, tt.userID)
				mutex.Unlock()
			}
		})
	}
}
