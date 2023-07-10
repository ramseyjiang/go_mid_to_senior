package nonblock

import (
	"fmt"
	"testing"
)

func TestSelectPattern(t *testing.T) {
	messages := make(chan string, 1)
	signals := make(chan bool, 1)

	// The select statement tries to send a value "hi" to the messages channel.
	t.Run("Non-blocking send", func(t *testing.T) {
		select {
		case messages <- "hi":
			msg := <-messages
			fmt.Printf("message send %v\n", msg)
			if msg != "hi" {
				t.Errorf("Expected 'hi', got %s", msg)
			}
		default:
			fmt.Println("no message sent")
			t.Error("Expected to send a message, but couldn't")
		}
	})

	t.Run("Non-blocking receive", func(t *testing.T) {
		// It sends a value "hi" to the messages channel at first.
		messages <- "hi"
		// The select statement tries to receive a value from the messages channel.
		select {
		case msg := <-messages:
			fmt.Printf("message receive %v\n", msg)
			if msg != "hi" {
				t.Errorf("Expected 'hi', got %s", msg)
			}
		default:
			fmt.Println("no message sent")
			t.Error("Expected to receive a message, but didn't")
		}
	})

	t.Run("Non-blocking multi-way select", func(t *testing.T) {
		messages <- "hi"
		signals <- true
		select {
		case msg := <-messages:
			fmt.Printf("message receive %v\n", msg)
			if msg != "hi" {
				t.Errorf("Expected 'hi', got %s", msg)
			}
		case sig := <-signals:
			fmt.Printf("sig receive %v\n", sig)
			if sig != true {
				t.Errorf("Expected 'true', got %v", sig)
			}
		default:
			fmt.Println("no message sent")
			t.Error("Expected to receive either a message or a signal, but didn't")
		}
	})
}
