package dynamicwp

import (
	"testing"
)

func TestWorkerPool(t *testing.T) {
	t.Run("Test Task Processing", func(t *testing.T) {
		wp := NewWorkerPool(5)
		go wp.Start()

		tasks := []Task{
			{"player1", "move"},
			{"player2", "jump"},
			{"player3", "shoot"},
		}

		for _, task := range tasks {
			wp.Tasks <- task
		}

		wp.Stop()
	})
}
