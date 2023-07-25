package tasks

import (
	"sync"
	"testing"
	"time"
)

func TestTaskSystem(t *testing.T) {
	t.Run("two workers", func(t *testing.T) {
		tasks := make(chan Task)
		done := make(chan bool)
		var wg sync.WaitGroup

		start := time.Now()

		go taskProducer(tasks, done)

		for i := 1; i <= 2; i++ {
			wg.Add(1)
			go worker(i, tasks, done, &wg)
		}

		<-done
		wg.Wait()

		elapsed := time.Since(start)
		t.Logf("Time taken for 3 workers: %s", elapsed)
	})

	t.Run("three workers", func(t *testing.T) {
		tasks := make(chan Task)
		done := make(chan bool)
		var wg sync.WaitGroup

		start := time.Now()

		go taskProducer(tasks, done)

		for i := 1; i <= 3; i++ {
			wg.Add(1)
			go worker(i, tasks, done, &wg)
		}

		<-done
		wg.Wait()

		elapsed := time.Since(start)
		t.Logf("Time taken for 3 workers: %s", elapsed)
	})

	t.Run("five workers", func(t *testing.T) {
		tasks := make(chan Task)
		done := make(chan bool)
		var wg sync.WaitGroup

		start := time.Now()

		go taskProducer(tasks, done)

		for i := 1; i <= 5; i++ {
			wg.Add(1)
			go worker(i, tasks, done, &wg)
		}

		<-done
		wg.Wait()

		elapsed := time.Since(start)
		t.Logf("Time taken for 3 workers: %s", elapsed)
	})
}
