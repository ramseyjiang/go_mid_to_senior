package callcenter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCallRateLimitMiddleware(t *testing.T) {
	lb := NewLeakyBucket(2, 50*time.Millisecond)
	lb.StartLeaking()
	defer lb.StopLeaking()

	tests := []struct {
		name              string
		numberOfCalls     int
		delayBetweenCalls time.Duration
		expectedStatus    []int
	}{
		{"HandleAllCalls", 2, 0, []int{http.StatusOK, http.StatusOK}},
		{"ExceedCapacity", 5, 60 * time.Millisecond, []int{http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK}},
		{"AfterProcessing", 1, 60 * time.Millisecond, []int{http.StatusOK}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := CallRateLimitMiddleware(lb, http.HandlerFunc(CallHandler))

			for i := 0; i < tt.numberOfCalls; i++ {
				req := httptest.NewRequest("GET", "/call", nil)
				rr := httptest.NewRecorder()

				handler.ServeHTTP(rr, req)
				fmt.Println(tt.expectedStatus[i])
				// Check the response status
				if status := rr.Code; status != tt.expectedStatus[i] {
					t.Errorf("%s: call %d returned wrong status code: got %v want %v",
						tt.name, i+1, status, tt.expectedStatus[i])
				}

				// Introduce a delay between calls
				time.Sleep(tt.delayBetweenCalls)
			}
		})
	}
}
