package apictrl

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimitMiddleware(t *testing.T) {
	tests := []struct {
		name             string
		numberOfRequests int
		delay            time.Duration
		expectedStatus   int
	}{
		{"WithinLimit", 3, 0, http.StatusOK},
		{"ExceedLimit", 6, 0, http.StatusTooManyRequests},
		{"AfterRefill", 1, 1 * time.Second, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bucket := NewTokenBucket(capacity, refillRate) // Reset bucket for each test

			handler := RateLimitMiddleware(bucket, http.HandlerFunc(APIHandler))

			for i := 0; i < tt.numberOfRequests; i++ {
				req := httptest.NewRequest("GET", "/api", nil)
				rr := httptest.NewRecorder()

				handler.ServeHTTP(rr, req)

				if i == tt.numberOfRequests-1 {
					if status := rr.Code; status != tt.expectedStatus {
						t.Errorf("%s: handler returned wrong status code: got %v want %v",
							tt.name, status, tt.expectedStatus)
					}
				}

				time.Sleep(tt.delay) // Delay for the next request
			}
		})
	}
}
