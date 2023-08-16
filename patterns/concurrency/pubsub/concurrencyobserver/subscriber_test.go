package concurrencyobserver

import (
	"bytes"
	"testing"
	"time"
)

func TestWriter(t *testing.T) {
	t.Run("Notify", func(t *testing.T) {
		var buf bytes.Buffer
		sub := NewWriterSubscriber(1, &buf)

		err := sub.Notify("Test")
		if err != nil {
			t.Fatalf("failed to notify: %v", err)
		}

		time.Sleep(50 * time.Millisecond) // Allow some time for the message to be written

		if buf.String() != "(W1): Test\n" {
			t.Fatalf("expected '(W1): Test\n' but got %v", buf.String())
		}
	})

	t.Run("Timeout", func(t *testing.T) {
		var buf bytes.Buffer
		sub := NewWriterSubscriber(1, &buf)

		close(sub.(*writerSubscriber).in) // Close the channel to induce a timeout

		err := sub.Notify("Test")
		if err == nil {
			t.Fatal("expected a timeout error but got none")
		}
	})

	t.Run("Close", func(t *testing.T) {
		var buf bytes.Buffer
		sub := NewWriterSubscriber(1, &buf)

		sub.Notify("Before Close")
		sub.Close()

		time.Sleep(50 * time.Millisecond) // Allow some time for the message to be written

		if buf.String() != "(W1): Before Close\n" {
			t.Fatalf("expected '(W1): Before Close\n' but got %v", buf.String())
		}

		err := sub.Notify("After Close")
		if err == nil || err.Error() != "send on closed channel" {
			t.Fatal("expected a 'send on closed channel' error when notifying after close, got:", err)
		}
	})
}
