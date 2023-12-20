package cdntraffic

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type TokenBucket struct {
	Capacity   int
	Tokens     int
	LastRefill time.Time
	RefillRate int // Tokens added per minute
	Mutex      sync.Mutex
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		Capacity:   capacity,
		Tokens:     capacity, // Start with a full bucket
		LastRefill: time.Now(),
		RefillRate: refillRate,
	}
}

func (tb *TokenBucket) Refill() {
	tb.Mutex.Lock()
	defer tb.Mutex.Unlock()

	now := time.Now()
	duration := now.Sub(tb.LastRefill)
	refillTokens := int(duration.Minutes()) * tb.RefillRate

	tb.Tokens += refillTokens
	if tb.Tokens > tb.Capacity {
		tb.Tokens = tb.Capacity
	}
	tb.LastRefill = now
}

func (tb *TokenBucket) AllowRequest() bool {
	tb.Refill()

	if tb.Tokens > 0 {
		tb.Tokens--
		return true
	}
	return false
}

func BandwidthLimitMiddleware(tb *TokenBucket, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !tb.AllowRequest() {
			http.Error(w, "Bandwidth limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Content delivered successfully")
}
