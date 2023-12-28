package callcenter

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LeakyBucket struct {
	Capacity int
	Queue    chan bool
	Rate     time.Duration
	wg       sync.WaitGroup
}

func NewLeakyBucket(capacity int, rate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		Capacity: capacity,
		Queue:    make(chan bool, capacity),
		Rate:     rate,
	}
}

func (lb *LeakyBucket) AddCall() bool {
	select {
	case lb.Queue <- true:
		return true
	default:
		return false // Bucket is full
	}
}

func (lb *LeakyBucket) StartLeaking() {
	lb.wg.Add(1)
	go func() {
		defer lb.wg.Done()
		for call := range lb.Queue {
			if !call {
				break // Stop leaking when a false value is received
			}
			time.Sleep(lb.Rate)
			// Process the call
		}
	}()
}

func (lb *LeakyBucket) StopLeaking() {
	lb.Queue <- false // Send a signal to stop processing
	lb.wg.Wait()
	close(lb.Queue)
}

func CallRateLimitMiddleware(lb *LeakyBucket, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !lb.AddCall() {
			http.Error(w, "Call limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CallHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Call handled successfully")
}
