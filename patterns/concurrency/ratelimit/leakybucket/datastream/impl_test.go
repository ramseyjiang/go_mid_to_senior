package datastream

import (
	"testing"
	"time"
)

func TestLeakyBucketDataStreaming(t *testing.T) {
	tests := []struct {
		name                string
		capacity            int
		rate                time.Duration
		numberOfPackets     int
		delayBetweenPackets time.Duration
		expectedSuccesses   int
	}{
		{"WithinCapacity", capacity, leakyRate, 5, 50 * time.Millisecond, 5},
		{"ExceedCapacity", capacity, leakyRate, 7, 10 * time.Millisecond, 5},
		{"AfterLeak", capacity, leakyRate, 6, 150 * time.Millisecond, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := NewLeakyBucket(tt.capacity, tt.rate)
			lb.StartStreaming()
			defer lb.StopStreaming()

			successes := 0
			for i := 0; i < tt.numberOfPackets; i++ {
				if lb.AddPacket(DataPacket{}) {
					successes++
				}
				time.Sleep(tt.delayBetweenPackets)
			}

			if successes != tt.expectedSuccesses {
				t.Errorf("%s: expected %d successful packets, but got %d", tt.name, tt.expectedSuccesses, successes)
			}
		})
	}
}
