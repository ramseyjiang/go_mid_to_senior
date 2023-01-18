package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ramseyjiang/go_mid_to_senior/puzzles/ratelimit"
)

// When this file is running, the first 3 token requests are processed immediately,
// then once the limit is reached each subsequent request has to wait until a token becomes available.
// At any time we never have more than 3 tokens in use, which is exactly what we want!

func main() {
	r, err := ratelimit.NewTokenBucketRateLimiter(&ratelimit.Config{
		Limit: 5,
		// Reset tokens manually after 10 seconds
		TokenResetsAfter: 10 * time.Second,
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
