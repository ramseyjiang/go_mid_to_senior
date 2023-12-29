package cloudprovider

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	capacity  = 2
	leakyRate = 50 * time.Millisecond
)

type LeakyBucket struct {
	Capacity int
	Queue    chan bool
	Rate     time.Duration
	wg       sync.WaitGroup
	stop     chan struct{}
}

func NewLeakyBucket(capacity int, rate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		Capacity: capacity,
		Queue:    make(chan bool, capacity),
		Rate:     rate,
		stop:     make(chan struct{}),
	}
}

func (lb *LeakyBucket) RequestResource() bool {
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
		ticker := time.NewTicker(lb.Rate)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if len(lb.Queue) > 0 {
					<-lb.Queue // Process a resource request
				}
			case <-lb.stop:
				return
			}
		}
	}()
}

func (lb *LeakyBucket) StopLeaking() {
	close(lb.stop)
	lb.wg.Wait()
}

func RequestRateMiddleware(lb *LeakyBucket, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !lb.RequestResource() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Resource request processed successfully")
	// do real requests logic
}
