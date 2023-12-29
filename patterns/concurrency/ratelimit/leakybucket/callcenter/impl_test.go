package callcenter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCallRateLimitMiddleware(t *testing.T) {
	tests := []struct {
		name              string
		capacity          int
		rate              time.Duration
		numberOfCalls     int
		delayBetweenCalls time.Duration
		expectedStatus    []int
	}{
		{"WithinCapacity", capacity, leakyRate, 2, 0, []int{http.StatusOK, http.StatusOK}},
		{"ExceedCapacity", capacity, leakyRate, 3, 10 * time.Millisecond, []int{http.StatusOK, http.StatusOK, http.StatusTooManyRequests}},
		{"AfterLeak", capacity, leakyRate, 3, 60 * time.Millisecond, []int{http.StatusOK, http.StatusOK, http.StatusOK}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := NewLeakyBucket(tt.capacity, tt.rate)
			lb.StartLeaking()
			defer lb.StopLeaking()

			handler := CallRateLimitMiddleware(lb, http.HandlerFunc(CallHandler))

			for i := 0; i < tt.numberOfCalls; i++ {
				req := httptest.NewRequest("GET", "/call", nil)
				rr := httptest.NewRecorder()

				handler.ServeHTTP(rr, req)

				if status := rr.Code; status != tt.expectedStatus[i] {
					t.Errorf("%s: call %d returned wrong status code: got %v want %v",
						tt.name, i+1, status, tt.expectedStatus[i])
				}

				time.Sleep(tt.delayBetweenCalls)
			}
		})
	}
}
