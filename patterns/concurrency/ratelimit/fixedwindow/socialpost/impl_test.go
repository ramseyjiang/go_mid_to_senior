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
		resetAfter        time.Duration
	}{
		{"UnderLimit", "user1", 5, false, http.StatusOK, 0},
		{"AtLimit", "user2", limitNumber, false, http.StatusOK, 0},
		{"OverLimit", "user3", limitNumber, true, http.StatusTooManyRequests, 0},
		{"ResetAfterOneMinute", "user4", limitNumber, true, http.StatusOK, 1 * time.Minute},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock request
			req := httptest.NewRequest("POST", "/comment", nil)
			req.Header.Set("X-User-ID", tt.userID)

			handler := RateLimitMiddleware(http.HandlerFunc(CommentHandler))

			// Send requests up to the limit
			for i := 0; i < tt.requestCount; i++ {
				rr := httptest.NewRecorder()
				handler.ServeHTTP(rr, req)
			}

			// Wait for reset interval if needed
			if tt.resetAfter > 0 {
				time.Sleep(tt.resetAfter)
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
