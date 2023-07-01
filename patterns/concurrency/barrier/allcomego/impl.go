package allcomego

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const amount = 5

type Barrier struct {
	total int
	count int
	mutex sync.Mutex
	cond  *sync.Cond
}

func NewBarrier(count int) *Barrier {
	b := &Barrier{total: count, count: 0}
	b.cond = sync.NewCond(&b.mutex)
	return b
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count++
	if b.count >= b.total {
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.mutex.Unlock()
}

func student(id int, b *Barrier, wg *sync.WaitGroup) {
	defer wg.Done()
	// Each student takes a random amount of time to arrive
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("Student %d arrived at the meeting point.\n", id)
	b.Wait()
	fmt.Printf("Student %d starts the camping trip.\n", id)
}
