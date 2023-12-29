package cloudprovider

import (
	"testing"
	"time"
)

func TestResourceRateLimit(t *testing.T) {
	tests := []struct {
		name              string
		numberOfRequests  int
		delay             time.Duration
		expectedSuccesses int
	}{
		{"WithinCapacity", 2, 0, 2},
		{"ExceedCapacity", 4, 10 * time.Millisecond, 2},
		{"AfterLeak", 3, 150 * time.Millisecond, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := NewLeakyBucket(capacity, leakyRate)
			lb.StartLeaking()
			defer lb.StopLeaking()

			successes := 0
			for i := 0; i < tt.numberOfRequests; i++ {
				if lb.RequestResource() {
					successes++
				}
				time.Sleep(tt.delay)
			}

			if successes != tt.expectedSuccesses {
				t.Errorf("%s: expected %d successes, but got %d", tt.name, tt.expectedSuccesses, successes)
			}
		})
	}
}
