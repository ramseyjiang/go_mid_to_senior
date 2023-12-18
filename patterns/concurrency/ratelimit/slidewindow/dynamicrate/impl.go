package dynamicrate

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Define the const depends on your real environment
const (
	limitLoad            = 50
	lowerLoadHigherLimit = 100
	higherLoadLowerLimit = 50
)

type RequestWindow struct {
	Requests []time.Time
	Mutex    sync.Mutex
}

func NewRequestWindow() *RequestWindow {
	return &RequestWindow{
		Requests: []time.Time{},
	}
}

func (rw *RequestWindow) AddRequest(t time.Time) {
	rw.Mutex.Lock()
	defer rw.Mutex.Unlock()
	rw.Requests = append(rw.Requests, t)
}

func (rw *RequestWindow) AllowRequest(currentLoad int) bool {
	rw.Mutex.Lock()
	defer rw.Mutex.Unlock()

	now := time.Now()
	windowStart := now.Add(-1 * time.Minute)

	// Remove outdated requests
	validRequests := []time.Time{}
	for _, reqTime := range rw.Requests {
		if reqTime.After(windowStart) {
			validRequests = append(validRequests, reqTime)
		}
	}
	rw.Requests = validRequests

	// Dynamic rate limit based on current load
	var limit int
	if currentLoad < limitLoad {
		limit = lowerLoadHigherLimit // Less load, higher limit
	} else {
		limit = higherLoadLowerLimit // More load, lower limit
	}

	return len(rw.Requests) < limit
}

func RateLimitMiddleware(rw *RequestWindow, getCurrentLoad func() int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rw.AllowRequest(getCurrentLoad()) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		rw.AddRequest(time.Now())
		next.ServeHTTP(w, r)
	})
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request processed successfully")
	// Do some real logic
}
