package messagequeue

import (
	"fmt"
	"testing"
)

func TestMessageQueue(t *testing.T) {
	mq := NewMessageQueue()

	t.Run("Send and Receive", func(t *testing.T) {
		mq.Send("Message 1")
		mq.Send("Message 2")

		fmt.Println("Received:", mq.Receive())
		fmt.Println("Received:", mq.Receive())
	})
}
