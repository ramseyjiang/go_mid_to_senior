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

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")

		mutex.Lock()
		activity, exists := userActivities[userID]
		if !exists || time.Now().After(activity.ResetTime) {
			userActivities[userID] = &UserActivity{Count: 1, ResetTime: time.Now().Add(1 * time.Minute)}
			mutex.Unlock()
			next.ServeHTTP(w, r)
			return
		}

		if activity.Count >= 30 {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			mutex.Unlock()
			return
		}

		activity.Count++
		mutex.Unlock()
		next.ServeHTTP(w, r)
	})
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the post/comment creation
	fmt.Fprintf(w, "Post/comment created successfully")
}
