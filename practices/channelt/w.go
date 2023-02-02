package main

import (
	"fmt"
	"sync"
	"time"
)

const total = 10000000

func worker(jobs <-chan int, results chan<- int) {
	for {
		select {
		case j := <-jobs:
			switch j % 3 {
			case 0:
				j = j * 1
				results <- j
			case 1:
				j = j * 2
				results <- j
			case 2:
				j *= 3
				results <- j
			}
		default:
			return
		}
	}
}

func main() {
	jobs := make(chan int, total)
	results := make(chan int, total)

	var wg sync.WaitGroup
	wg.Add(total)
	for i := 1; i <= total; i++ {
		go func(i int) {
			defer wg.Done()
			jobs <- i
		}(i)
	}
	wg.Wait()
	defer close(jobs)

	for w := 1; w < 1000; w++ {
		go worker(jobs, results)
	}
	defer close(results)

	var sum uint64 = 0
	go func() {
		for {
			select {
			case res := <-results:
				sum += uint64(res)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(sum)
}
