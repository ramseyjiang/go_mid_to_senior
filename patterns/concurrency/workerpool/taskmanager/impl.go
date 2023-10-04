package taskmanager

import (
	"fmt"
	"sync"
	"time"
)

// Step 1: Define the Task
type Task struct {
	ID      int
	Execute func() error
}

func Worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	// Create a Worker Pool with workers
	for task := range tasks {
		fmt.Printf("Worker %d executing task %d\n", id, task.ID)
		err := task.Execute()
		if err != nil {
			return
		}
		wg.Done()
	}
}

func ScheduleTasks(workerCount, taskCount int) {
	var wg sync.WaitGroup
	tasks := make(chan Task, 10)

	// Step 3: Dispatch Tasks
	for i := 1; i <= workerCount; i++ {
		go Worker(i, tasks, &wg)
	}

	for i := 1; i <= taskCount; i++ {
		wg.Add(1)
		tasks <- Task{
			ID: i,
			Execute: func() error {
				time.Sleep(1 * time.Second)
				fmt.Println("Task executed")
				return nil
			},
		}
	}

	wg.Wait()    // step 4: Wait for Completion
	close(tasks) // step 5: Graceful Shutdown
}
