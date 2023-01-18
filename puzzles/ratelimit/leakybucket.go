package ratelimit

import (
	"errors"
	"time"
)

// NewLeakyBucketRateLimiter is used to for a solution
// that behind a throttler is to handle burst programs by only allowing a certain number of requests to be processed per time duration.
func NewLeakyBucketRateLimiter(conf *Config) (RateLimiter, error) {
	if conf.Throttle == 0 {
		return nil, errors.New("leaky bucket algorithm duration must be greater than zero")
	}

	m := NewManager(conf)

	// Throttle Await Function, it will loop over the ticker channel and block while waiting to receive a message from the in channel.
	await := func(throttle time.Duration) {
		// a time Ticker is used to schedule ticks at a specified interval.
		// The time ticker at here, I can synchronize the throttler to any allow rate limits token per tick.
		ticker := time.NewTicker(throttle)
		go func() {
			<-m.InChan
			m.TryGenerateToken()
			for {
				select {
				// Once a message is received it’ll call tryGenerateToken() and continue looping to wait for the next ticker.
				case <-m.InChan:
					<-ticker.C
					m.TryGenerateToken()
				case t := <-m.ReleaseChan:
					m.ReleaseToken(t)
				}
			}
		}()
	}

	// Call awaits to start
	await(conf.Throttle)
	return m, nil
}
