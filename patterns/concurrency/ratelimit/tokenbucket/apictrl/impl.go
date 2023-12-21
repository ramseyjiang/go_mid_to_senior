package apictrl

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	capacity   = 5
	refillRate = 1
)

type TokenBucket struct {
	Capacity   int
	Tokens     int
	RefillRate int // Tokens added per second
	Mutex      sync.Mutex
	LastRefill time.Time
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		Capacity:   capacity,
		Tokens:     capacity, // Start with a full bucket
		RefillRate: refillRate,
		LastRefill: time.Now(),
	}
}

func (tb *TokenBucket) Refill() {
	tb.Mutex.Lock()
	defer tb.Mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.LastRefill).Seconds()
	refillTokens := int(elapsed) * tb.RefillRate

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

func RateLimitMiddleware(tb *TokenBucket, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !tb.AllowRequest() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API request successful")

	// Do some real logic
}
