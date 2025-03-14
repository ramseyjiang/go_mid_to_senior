package socialpost

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type UserActivity struct {
	Count     int
	ResetTime time.Time
}

var (
	userActivities = make(map[string]*UserActivity)
	mutex          = &sync.Mutex{}
)

const limitNumber = 10

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")

		mutex.Lock()
		activity, exists := userActivities[userID]
		currentTime := time.Now()

		if !exists || currentTime.After(activity.ResetTime) {
			userActivities[userID] = &UserActivity{Count: 1, ResetTime: currentTime.Add(1 * time.Minute)}
			mutex.Unlock()
			next.ServeHTTP(w, r)
			return
		}

		if activity.Count >= limitNumber {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			mutex.Unlock()
			return
		}

		activity.Count++
		mutex.Unlock()
		next.ServeHTTP(w, r)
	})
}

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Comment created successfully")
	// Do the exact logic in this handler
}
