package ratelimit

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu          sync.Mutex
	requests    map[int]int
	timestamps  map[int]time.Time
	maxRequests int
	windowSize  time.Duration
	currentTime time.Time // Add currentTime to the struct
}

func New(maxRequests int, windowSizeSeconds int) *RateLimiter {
	return &RateLimiter{
		requests:    make(map[int]int),
		timestamps:  make(map[int]time.Time),
		maxRequests: maxRequests,
		windowSize:  time.Duration(windowSizeSeconds) * time.Second,
		currentTime: time.Now(), // Initialize currentTime
	}
}

// setCurrentTime is a helper method for testing. It sets the current time.
func (r *RateLimiter) setCurrentTime(t time.Time) {
	r.currentTime = t
}

// RateLimit checks if a request is allowed for the given customerId
func (r *RateLimiter) RateLimit(customerID int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := r.currentTime // Use the currentTime field

	// Check and reset the window if needed
	if timestamp, exists := r.timestamps[customerID]; !exists || now.Sub(timestamp) >= r.windowSize {
		r.timestamps[customerID] = now
		r.requests[customerID] = 0
	}

	// Allow or deny the request
	if r.requests[customerID] < r.maxRequests {
		r.requests[customerID]++
		return true
	}

	return false
}
