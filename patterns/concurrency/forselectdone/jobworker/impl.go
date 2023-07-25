package jobworker

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(id int, jobs <-chan Job, done <-chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case job, ok := <-jobs:
			if ok {
				fmt.Printf("Worker %d processing job %d\n", id, job.ID)
				time.Sleep(1 * time.Second) // simulate job processing time
			} else {
				wg.Done()
				return
			}
		case <-done:
			wg.Done()
			return
		}
	}
}
