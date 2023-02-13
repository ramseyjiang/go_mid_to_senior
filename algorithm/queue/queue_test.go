package queue

import (
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestQueue(t *testing.T) {
	var queue Queue

	// queue rules is the first in first out.
	queue.Enqueue(5)
	queue.Enqueue(4)
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)

	log.Println("Queue:", queue.Dump())

	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 5, queue.Dequeue())

	log.Println("Queue:", queue.Dump())

	assert.Equal(t, false, queue.IsEmpty())

	queue.Reset()

	assert.Equal(t, true, queue.IsEmpty())
}
