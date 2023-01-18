package ratelimit

import (
	"time"

	"github.com/pkg/errors"
)

// FixedWindowInterval represents a fixed window of time with start time, end time and an interval duration.
// The token expires time is the end time of current fixed window interval
type FixedWindowInterval struct {
	startTime time.Time
	endTime   time.Time
	interval  time.Duration
}

func (w *FixedWindowInterval) setWindowTime() {
	w.startTime = time.Now().UTC()
	w.endTime = time.Now().UTC().Add(w.interval)
}

// run is used to start a ticker of interval n that sets the fixed window time and calls a callback function per each interval.
func (w *FixedWindowInterval) run(cb func()) {
	go func() {
		ticker := time.NewTicker(w.interval)
		w.setWindowTime()
		for range ticker.C {
			cb()
			w.setWindowTime()
		}
	}()
}

// NewFixedWindowRateLimiter defines a FixedWindowInterval and override manager.
// makeToken to set the token ExpiresAt to be the fixed window end time.
// After that, it calls window.run and pass manager.releaseExpiredTokens as the callback.
func NewFixedWindowRateLimiter(conf *Config) (RateLimiter, error) {
	if conf.FixedInterval == 0 {
		return nil, errors.New("Interval must be greater than zero")
	}

	if conf.Limit == 0 {
		return nil, errors.New("Limit must be greater than zero")
	}

	m := NewManager(conf)
	w := &FixedWindowInterval{interval: conf.FixedInterval}

	// override the manager makeToken function
	m.makeToken = func() *Token {
		t := NewToken()
		t.ExpiresAt = w.endTime
		return t
	}

	await := func() {
		go func() {
			for {
				select {
				case <-m.InChan:
					m.TryGenerateToken()
				case token := <-m.ReleaseChan:
					m.ReleaseToken(token)
				}
			}
		}()
	}

	w.run(m.releaseExpiredTokens)
	await()
	return m, nil
}
