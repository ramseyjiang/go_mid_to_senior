package cloudprovider

import (
	"sync"
	"time"
)

const (
	capacity  = 2
	leakyRate = 100 * time.Millisecond
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
