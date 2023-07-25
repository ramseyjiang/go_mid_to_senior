package jobworker

import (
	"sync"
	"testing"
)

func TestJobProcessing(t *testing.T) {
	testCases := []struct {
		name       string
		numWorkers int
		numJobs    int
	}{
		{"1 worker, 5 jobs", 1, 5},
		{"2 workers, 5 jobs", 2, 5},
		{"3 workers, 5 jobs", 3, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jobs := make(chan Job)
			done := make(chan bool)
			var wg sync.WaitGroup

			for i := 1; i <= tc.numWorkers; i++ {
				wg.Add(1)
				go worker(i, jobs, done, &wg)
			}

			go func() {
				for i := 1; i <= tc.numJobs; i++ {
					jobs <- Job{ID: i}
				}
				close(jobs)
				done <- true
			}()

			wg.Wait()
		})
	}
}
