package main

import (
	"fmt"
	"time"
)

type semIFace interface {
	Acquire()
	Release()
}

type semaphore struct {
	semChan chan struct{}
}

const testMaxCurrency int = 3
const totalProcess int = 10
const minDoneNum int = 1
const timeFormat string = "2006-01-02 15:04:05"

func New(maxConcurrency int) semIFace {
	return &semaphore{
		semChan: make(chan struct{}, maxConcurrency),
	}
}

// Acquire method is used to lock resources and will be called before calling our heavy long-running process.
func (s *semaphore) Acquire() {
	s.semChan <- struct{}{}
}

// Release method should be called after the long-running process has been processed.
func (s *semaphore) Release() {
	<-s.semChan
}

// When we call Acquire(), the channel will be filled with an empty struct, this channel will be blocking if it reaches its maximum value.
// And when we call Release(), we take out the empty struct from the channel,
// and the channel will be available for the next value and the channel will be unblocking.
func main() {
	sem := New(testMaxCurrency)             // instantiate a new Semaphore by the size 3, that means maximum concurrent process will be limited to 3
	doneChan := make(chan bool, minDoneNum) // unbuffered done channel

	for i := 1; i <= totalProcess; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			longRunningProcess(v)
			if v == totalProcess {
				doneChan <- true
			}
		}(i)
	}
	<-doneChan
}

// longRunningProcess emulate some heavy tasks on longRunningProcess function (we use 2 seconds sleep time to block the process)
func longRunningProcess(taskID int) {
	fmt.Println(time.Now().Format(timeFormat), "Running task with ID", taskID)
	time.Sleep(2 * time.Second)
}

/**
The example above, our function can run 3 processes concurrently and since we use Semaphore,
our example needs 6 seconds to complete the process, in contrast without Semaphore.
By using Semaphore we can control access to a shared resource eg: Database, Network, Disk, etc.
*/
