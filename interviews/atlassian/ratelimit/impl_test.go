package ratelimit

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	tests := []struct {
		name            string
		maxRequests     int
		windowSize      int
		customerID      int
		requests        []int // Times in seconds when the requests are made
		expectedResults []bool
	}{
		{
			name:            "WithinCapacity",
			maxRequests:     3,
			windowSize:      10,
			customerID:      1,
			requests:        []int{1, 2, 3}, // Within the same window
			expectedResults: []bool{true, true, true},
		},
		{
			name:            "ExceedCapacity",
			maxRequests:     3,
			windowSize:      10,
			customerID:      1,
			requests:        []int{1, 2, 3, 4}, // Exceeds the capacity in the same window
			expectedResults: []bool{true, true, true, false},
		},
		{
			name:            "ResetAfterWindow",
			maxRequests:     3,
			windowSize:      10,
			customerID:      2,
			requests:        []int{1, 2, 12, 13}, // Reset after the window
			expectedResults: []bool{true, true, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter := New(tt.maxRequests, tt.windowSize)
			startTime := time.Now()

			for i, seconds := range tt.requests {
				currentTime := startTime.Add(time.Duration(seconds) * time.Second)
				limiter.setCurrentTime(currentTime)

				got := limiter.RateLimit(tt.customerID)
				if got != tt.expectedResults[i] {
					t.Errorf("Request at second %d: RateLimit() = %v, want %v", seconds, got, tt.expectedResults[i])
				}
			}
		})
	}
}
