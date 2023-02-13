package queue

import (
	"sync"
)

type Item interface{}

type Queue struct {
	items []Item
	mutex sync.Mutex
}

func (queue *Queue) Enqueue(item Item) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.items = append(queue.items, item)
}

func (queue *Queue) Dequeue() Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if len(queue.items) == 0 {
		return nil
	}

	lastItem := queue.items[0]
	queue.items = queue.items[1:]

	return lastItem
}

func (queue *Queue) IsEmpty() bool {
	if len(queue.items) == 0 {
		return true
	}
	return false
}

func (queue *Queue) Peek() Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if len(queue.items) == 0 {
		return nil
	}

	return queue.items[len(queue.items)-1]
}

func (queue *Queue) Reset() {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.items = nil
}

func (queue *Queue) Dump() []Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	var copiedQueue = make([]Item, len(queue.items))
	copy(copiedQueue, queue.items)

	return copiedQueue
}
