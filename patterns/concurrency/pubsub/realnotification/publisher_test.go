package realnotification

import (
	"bytes"
	"testing"
	"time"
)

func TestPublisher(t *testing.T) {
	t.Run("AddSubscriber", func(t *testing.T) {
		var buf bytes.Buffer
		pub := NewPublisher()
		sub := NewWriterSubscriber(1, &buf)
		go pub.Start()

		pub.AddSubscriberCh() <- sub
		time.Sleep(50 * time.Millisecond) // Allow some time for the subscriber to be added

		pub.PublishingCh() <- "Hello"
		time.Sleep(50 * time.Millisecond) // Allow some time for the message to be processed

		if buf.String() != "(W1): Hello\n" {
			t.Fatalf("expected '(W1): Hello\n' but got %v", buf.String())
		}

		pub.Stop()
	})

	t.Run("RemoveSubscriber", func(t *testing.T) {
		var buf bytes.Buffer
		pub := NewPublisher()
		sub := NewWriterSubscriber(1, &buf)
		go pub.Start()

		pub.AddSubscriberCh() <- sub
		time.Sleep(50 * time.Millisecond) // Allow some time for the subscriber to be added

		pub.RemoveSubscriberCh() <- sub
		time.Sleep(50 * time.Millisecond) // Allow some time for the subscriber to be removed

		pub.PublishingCh() <- "Hello"
		time.Sleep(50 * time.Millisecond) // Allow some time to ensure the message is not received

		if buf.String() != "" {
			t.Fatalf("expected no messages but got %v", buf.String())
		}

		pub.Stop()
	})
}
