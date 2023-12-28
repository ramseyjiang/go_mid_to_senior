package callcenter

import (
	"fmt"
	"net/http"
	"time"
)

const (
	capacity = 2
	callRate = 50 * time.Millisecond
)

type LeakyBucket struct {
	Capacity int
	Queue    chan bool
	Rate     time.Duration
}

func NewLeakyBucket(capacity int, rate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		Capacity: capacity,
		Queue:    make(chan bool, capacity),
		Rate:     rate,
	}
}

func (lb *LeakyBucket) AllowCall() bool {
	select {
	case lb.Queue <- true:
		return true
	default:
		return false
	}
}

func (lb *LeakyBucket) StartProcessing() {
	go func() {
		for range time.Tick(lb.Rate) {
			<-lb.Queue
		}
	}()
}

func CallRateLimitMiddleware(lb *LeakyBucket, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !lb.AllowCall() {
			http.Error(w, "Call limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func HandleCalls(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Call handled successfully")
	// Do more real logic
}
