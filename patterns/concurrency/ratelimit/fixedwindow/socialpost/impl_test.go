package socialpost

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimitMiddleware(t *testing.T) {
	tests := []struct {
		name              string
		userID            string
		requestCount      int
		additionalRequest bool
		expectedStatus    int
	}{
		{"UnderLimit", "user1", 20, false, http.StatusOK},
		{"AtLimit", "user2", 30, false, http.StatusOK},
		{"OverLimit", "user3", 30, true, http.StatusTooManyRequests},
		{"ResetAfterOneMinute", "user4", 30, true, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/post", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("X-User-ID", tt.userID)

			handler := RateLimitMiddleware(http.HandlerFunc(PostHandler))

			// Send requests up to the limit
			for i := 0; i < tt.requestCount; i++ {
				rr := httptest.NewRecorder()
				handler.ServeHTTP(rr, req)
			}

			// Wait for reset interval if needed
			if tt.name == "ResetAfterOneMinute" {
				time.Sleep(1 * time.Minute)
			}

			// Send an additional request if required
			if tt.additionalRequest {
				rr := httptest.NewRecorder()
				handler.ServeHTTP(rr, req)

				// Assert for the additional request
				if status := rr.Code; status != tt.expectedStatus {
					t.Errorf("handler returned wrong status code: got %v want %v",
						status, tt.expectedStatus)
				}
			}
		})
	}
}
