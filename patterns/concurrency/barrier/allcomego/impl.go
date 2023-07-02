package allcomego

import (
	"sync"
)

const amount = 5

type Barrier struct {
	total int
	count int
	mutex sync.Mutex
	cond  *sync.Cond
}

func NewBarrier(count int) *Barrier {
	b := &Barrier{
		total: count,
		count: 0,
	}
	b.cond = sync.NewCond(&b.mutex)
	return b
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count++
	if b.count >= b.total {
		b.count = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.mutex.Unlock()
}
