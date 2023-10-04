package taskmanager

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	tasks := make(chan Task, 10)
	var wg sync.WaitGroup

	t.Run("TestTaskExecution", func(t *testing.T) {
		go Worker(1, tasks, &wg)

		t.Run("SubTestTask1", func(t *testing.T) {
			wg.Add(1)
			tasks <- Task{
				ID: 1,
				Execute: func() error {
					time.Sleep(1 * time.Second)
					fmt.Println("Task 1 executed")
					return nil
				},
			}
		})

		t.Run("SubTestTask2", func(t *testing.T) {
			wg.Add(1)
			tasks <- Task{
				ID: 2,
				Execute: func() error {
					time.Sleep(1 * time.Second)
					fmt.Println("Task 2 executed")
					return nil
				},
			}
		})

		t.Run("TestSchedulingAndExecution", func(t *testing.T) {
			// Capture the standard output.
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Schedule and execute tasks.
			ScheduleTasks(3, 5)

			// Restore the standard output and read the captured output.
			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = rescueStdout

			// Validate the output.
			expectedSubstring := "Task executed"
			if !strings.Contains(string(out), expectedSubstring) {
				t.Errorf("Expected output to contain: %s, got: %s", expectedSubstring, string(out))
			}
		})
	})

	wg.Wait()
	close(tasks)
}
