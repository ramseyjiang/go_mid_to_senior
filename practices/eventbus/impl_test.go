package eventbus

import (
	"fmt"
	"testing"
)

func TestEventBus(t *testing.T) {
	eb := NewEventBus()

	t.Run("Subscribe and Publish", func(t *testing.T) {
		eb.Subscribe("event1", func(data string) {
			fmt.Println("Subscriber 1 received:", data)
		})
		eb.Subscribe("event1", func(data string) {
			fmt.Println("Subscriber 2 received:", data)
		})

		eb.Publish("event1", "Hello from event1!")
	})
}
