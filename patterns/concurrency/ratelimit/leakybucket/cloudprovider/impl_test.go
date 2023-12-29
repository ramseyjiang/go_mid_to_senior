package cloudprovider

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequestRateMiddleware(t *testing.T) {
	tests := []struct {
		name              string
		capacity          int
		rate              time.Duration
		numberOfRequests  int
		delayBetweenCalls time.Duration
		expectedStatus    []int
	}{
		{
			name:              "WithinCapacity",
			capacity:          2,
			rate:              100 * time.Millisecond,
			numberOfRequests:  2,
			delayBetweenCalls: 10 * time.Millisecond,
			expectedStatus:    []int{http.StatusOK, http.StatusOK},
		},
		{
			name:              "ExceedCapacity",
			capacity:          2,
			rate:              100 * time.Millisecond,
			numberOfRequests:  3,
			delayBetweenCalls: 10 * time.Millisecond,
			expectedStatus:    []int{http.StatusOK, http.StatusOK, http.StatusTooManyRequests},
		},
		{
			name:              "AfterLeak",
			capacity:          2,
			rate:              100 * time.Millisecond,
			numberOfRequests:  3,
			delayBetweenCalls: 150 * time.Millisecond,
			expectedStatus:    []int{http.StatusOK, http.StatusOK, http.StatusOK},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := NewLeakyBucket(tt.capacity, tt.rate)
			lb.StartLeaking()
			defer lb.StopLeaking()

			handler := RequestRateMiddleware(lb, http.HandlerFunc(RequestHandler))

			for i := 0; i < tt.numberOfRequests; i++ {
				req := httptest.NewRequest("GET", "/resource", nil)
				rr := httptest.NewRecorder()

				handler.ServeHTTP(rr, req)

				if status := rr.Code; status != tt.expectedStatus[i] {
					t.Errorf("%s: request %d returned wrong status code: got %v want %v",
						tt.name, i+1, status, tt.expectedStatus[i])
				}

				time.Sleep(tt.delayBetweenCalls)
			}
		})
	}
}
