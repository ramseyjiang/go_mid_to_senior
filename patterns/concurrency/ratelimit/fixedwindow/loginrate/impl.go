package loginrate

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LoginAttempt struct {
	Count     int
	ResetTime time.Time
}

var (
	loginAttempts = make(map[string]*LoginAttempt)
	mutex         = &sync.Mutex{}
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")

		mutex.Lock()
		attempt, exists := loginAttempts[username]
		currentTime := time.Now()

		if !exists || currentTime.After(attempt.ResetTime) {
			loginAttempts[username] = &LoginAttempt{Count: 1, ResetTime: currentTime.Add(15 * time.Minute)}
			mutex.Unlock()
			next.ServeHTTP(w, r)
			return
		}

		if attempt.Count >= 5 {
			http.Error(w, "Too many login attempts", http.StatusTooManyRequests)
			mutex.Unlock()
			return
		}

		attempt.Count++
		mutex.Unlock()
		next.ServeHTTP(w, r)
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the login logic
	fmt.Fprintf(w, "Login attempt successful")
}
