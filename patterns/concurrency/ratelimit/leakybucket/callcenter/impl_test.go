package callcenter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCallRateLimitMiddleware(t *testing.T) {
	lb := NewLeakyBucket(capacity, callRate)
	lb.StartProcessing()

	tests := []struct {
		name           string
		delay          time.Duration
		expectedStatus int
	}{
		{"FirstCall", 0, http.StatusOK},
		{"SecondCall", 0, http.StatusOK},
		{"ExceedLimit", 0, http.StatusTooManyRequests},
		{"AfterProcessing", 60 * time.Millisecond, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := CallRateLimitMiddleware(lb, http.HandlerFunc(HandleCalls))

			time.Sleep(tt.delay)
			req := httptest.NewRequest("GET", "/call", nil)
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("%s: handler returned wrong status code: got %v want %v",
					tt.name, status, tt.expectedStatus)
			}
		})
	}
}
