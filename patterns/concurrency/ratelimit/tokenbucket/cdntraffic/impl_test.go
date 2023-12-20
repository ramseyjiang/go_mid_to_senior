package cdntraffic

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBandwidthLimitMiddleware(t *testing.T) {
	tests := []struct {
		name             string
		numberOfRequests int
		expectedStatus   int
	}{
		{"ImmediateRequest", 1, http.StatusOK},
		{"ExceedLimit", 6, http.StatusTooManyRequests}, // 6 requests to exceed the limit of 5
		{"AfterRefill", 1, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bucket := NewTokenBucket(capacity, refillRate) // Reset bucket for each test

			handler := BandwidthLimitMiddleware(bucket, http.HandlerFunc(ContentHandler))

			var lastStatus int
			for i := 0; i < tt.numberOfRequests; i++ {
				req := httptest.NewRequest("GET", "/content", nil)
				rr := httptest.NewRecorder()
				handler.ServeHTTP(rr, req)
				lastStatus = rr.Code

				if tt.name == "AfterRefill" {
					time.Sleep(1 * time.Minute) // Wait for tokens to refill
				}
			}

			if lastStatus != tt.expectedStatus {
				t.Errorf("%s: handler returned wrong status code: got %v want %v", tt.name, lastStatus, tt.expectedStatus)
			}
		})
	}
}
