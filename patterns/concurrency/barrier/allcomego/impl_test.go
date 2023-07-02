package allcomego

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestBarrier(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) // Ensure randomness
	var wg sync.WaitGroup
	barrier := NewBarrier(amount) // 5 students in total
	var mutex sync.Mutex
	var sum = 0

	for i := 1; i <= amount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Each student takes a random amount of time to arrive
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			fmt.Printf("Student %d arrived at the meeting point.\n", id)
			barrier.Wait()
			mutex.Lock()
			sum++
			mutex.Unlock()
			fmt.Printf("Student %d starts the camping trip.\n", id)
		}(i)
	}

	wg.Wait()

	// Check whether the last student who arrived is also the last one who started the trip
	if sum != amount {
		t.Errorf("Not all students arrive, we cannot go.")
	} else {
		fmt.Printf("All %d students arrive, let's go.", sum)
	}
}
