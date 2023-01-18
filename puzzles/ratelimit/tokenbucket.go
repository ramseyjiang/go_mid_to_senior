package ratelimit

import "github.com/pkg/errors"

// NewTokenBucketRateLimiter returns a max concurrency rate limiter
func NewTokenBucketRateLimiter(conf *Config) (RateLimiter, error) {
	if conf.Limit <= 0 {
		return nil, errors.New("ErrInvalidLimit")
	}

	m := NewManager(conf)

	// max concurrency await function
	await := func() {
		go func() {
			for {
				select {
				case <-m.InChan:
					m.TryGenerateToken()
				case t := <-m.ReleaseChan:
					m.ReleaseToken(t)
				}
			}
		}()
	}

	await()
	return m, nil
}
