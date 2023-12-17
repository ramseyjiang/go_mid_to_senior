package loginrate

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const maxLoginAttempts = 5

type LoginAttempt struct {
	Count     int
	ResetTime time.Time
}

type RateLimiter struct {
	loginAttempts map[string]*LoginAttempt
	mutex         sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		loginAttempts: make(map[string]*LoginAttempt),
	}
}

func (rl *RateLimiter) GetLoginAttempt(username string) *LoginAttempt {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	attempt, exists := rl.loginAttempts[username]
	if !exists || time.Now().After(attempt.ResetTime) {
		rl.loginAttempts[username] = &LoginAttempt{Count: 1, ResetTime: time.Now().Add(15 * time.Minute)}
		return rl.loginAttempts[username]
	}
	return attempt
}

func RateLimitMiddleware(rl *RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")

		attempt := rl.GetLoginAttempt(username)

		if attempt.Count > maxLoginAttempts {
			http.Error(w, "Too many login attempts", http.StatusTooManyRequests)
			return
		}

		attempt.Count++
		next.ServeHTTP(w, r)
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login attempt successful")
	// Add your exact logic after login
}
