package tasks

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

func taskProducer(tasks chan<- Task, done chan<- bool) {
	for i := 1; i <= 6; i++ {
		tasks <- Task{ID: i}

		// simulate time to produce a task
		time.Sleep(100 * time.Millisecond)
	}
	close(tasks)
	done <- true
}

func worker(id int, tasks <-chan Task, done <-chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case task, ok := <-tasks:
			if ok {
				fmt.Printf("Worker %d processing task %d\n", id, task.ID)

				// simulate time to process a task
				time.Sleep(500 * time.Millisecond)
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
