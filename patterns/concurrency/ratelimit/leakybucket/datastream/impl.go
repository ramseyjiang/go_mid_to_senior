package datastream

import (
	"sync"
	"time"
)

const (
	capacity  = 5
	leakyRate = 100 * time.Millisecond
)

type DataPacket struct {
	// Define the structure of your data packet
}

type LeakyBucket struct {
	Capacity int
	Queue    chan DataPacket
	Rate     time.Duration
	wg       sync.WaitGroup
	stop     chan struct{}
}

func NewLeakyBucket(capacity int, rate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		Capacity: capacity,
		Queue:    make(chan DataPacket, capacity),
		Rate:     rate,
		stop:     make(chan struct{}),
	}
}

func (lb *LeakyBucket) AddPacket(packet DataPacket) bool {
	select {
	case lb.Queue <- packet:
		return true
	default:
		return false // Bucket is full
	}
}

func (lb *LeakyBucket) StartStreaming() {
	lb.wg.Add(1)
	go func() {
		defer lb.wg.Done()
		ticker := time.NewTicker(lb.Rate)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if len(lb.Queue) > 0 {
					<-lb.Queue
					// Process the data packet, e.g., send it to a consumer
				}
			case <-lb.stop:
				return
			}
		}
	}()
}

func (lb *LeakyBucket) StopStreaming() {
	close(lb.stop)
	lb.wg.Wait()
}
