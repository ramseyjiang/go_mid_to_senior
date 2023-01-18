package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ramseyjiang/go_mid_to_senior/puzzles/ratelimit"
)

// run "go run main.go"
// the first 5 rate limit requests are granted.
// Then there is 15 seconds waiting period,
// After that, the next 5 are granted in the following fixed window interval.
func main() {
	r, err := ratelimit.NewFixedWindowRateLimiter(&ratelimit.Config{
		Limit:         5,
		FixedInterval: 15 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	doWork := func(id int) {
		// Acquire a rate limit token
		token, err := r.Acquire()
		fmt.Printf("Rate Limit Token %s acquired at %s...\n", token.ID, time.Now().UTC())
		if err != nil {
			panic(err)
		}
		// Simulate some work
		n := rand.Intn(5)
		fmt.Printf("Worker %d Sleeping %d seconds...\n", id, n)
		time.Sleep(time.Duration(n) * time.Second)
		fmt.Printf("Worker %d Done\n", id)
		r.Release(token)
		wg.Done()
	}

	// Spin up 10 workers that need a rate limit resource
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doWork(i)
	}

	wg.Wait()
}
