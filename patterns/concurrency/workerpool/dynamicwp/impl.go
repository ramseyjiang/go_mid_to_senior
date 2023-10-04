package dynamicwp

import (
	"sync"
	"time"
)

// Step 1: Define the Task
type Task struct {
	PlayerID string
	Action   string
}

type WorkerPool struct {
	Tasks       chan Task
	WorkerCount int
	QuitChan    chan bool
}

func NewWorkerPool(workerCount int) *WorkerPool {
	// Step 2: Create a Worker Pool with workers
	return &WorkerPool{
		Tasks:       make(chan Task, 100),
		WorkerCount: workerCount,
		QuitChan:    make(chan bool),
	}
}

func (wp *WorkerPool) Start() {
	var wg sync.WaitGroup

	// Step 3: Dispatch Tasks
	for i := 0; i < wp.WorkerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case task := <-wp.Tasks:
					// Simulate processing the task
					time.Sleep(100 * time.Millisecond)
					ProcessTask(task)
				case <-wp.QuitChan:
					return
				}
			}
		}()
	}

	wg.Wait() // step 4: Wait for Completion
}

func (wp *WorkerPool) Stop() {
	close(wp.QuitChan) // step 5: Graceful Shutdown
}

func ProcessTask(t Task) {
	// Process the task, e.g., update the game state based on player action
}
