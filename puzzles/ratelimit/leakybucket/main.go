package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ramseyjiang/go_mid_to_senior/puzzles/ratelimit"
)

// This example trys to create new NewLeakyBucketRateLimiter with
// a throttle duration of 1 second and generate 10 request workers that need a Rate Limit Token.
// Each worker receives a token throttled at 1 second intervals, as desired!

func main() {
	r, err := ratelimit.NewLeakyBucketRateLimiter(&ratelimit.Config{
		Throttle: 1 * time.Second,
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
		time.Sleep(time.Duration(n) * time.Second) // assume sleep rand.Intn(5) seconds
		fmt.Printf("Worker %d Done\n", id)
		wg.Done()
	}

	// Spin up 10 workers that need a rate limit resource
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doWork(i)
	}

	wg.Wait()
}
