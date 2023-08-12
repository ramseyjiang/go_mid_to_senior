package concurrencyobserver

import (
	"testing"
	"time"
)

type mockSubscriber struct {
	notifyTestingFunc func(msg interface{})
	closeTestingFunc  func()
}

func (m *mockSubscriber) Notify(msg interface{}) error {
	m.notifyTestingFunc(msg)
	return nil
}

func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}

func TestPublisher(t *testing.T) {
	pub := NewPublisher().(*publisher) // Assert to concrete type

	// Initialize the publisher's channels
	pub.addSubCh = make(chan Subscriber)
	pub.removeSubCh = make(chan Subscriber)
	pub.in = make(chan interface{})
	pub.stop = make(chan struct{})

	// Start the publisher in a goroutine
	go pub.start()

	msgToSend := "Hello"
	notified := make(chan bool, 1)
	closed := make(chan bool, 1)

	sub := &mockSubscriber{
		notifyTestingFunc: func(received interface{}) {
			if s, ok := received.(string); ok && s == msgToSend {
				notified <- true
			} else {
				notified <- false
			}
		},
		closeTestingFunc: func() {
			closed <- true
		},
	}

	// Add a subscriber
	go func() { pub.addSubCh <- sub }()
	time.Sleep(50 * time.Millisecond) // Give a short delay to ensure the subscriber is added

	// Publish a message
	go func() { pub.in <- msgToSend }()
	select {
	case success := <-notified:
		if !success {
			t.Error("Subscriber did not receive the expected message")
		}
	case <-time.After(1 * time.Second):
		t.Error("Timed out waiting for subscriber notification")
	}

	// Assert subscriber count
	if len(pub.subscribers) != 1 {
		t.Errorf("Expected 1 subscriber but got %d", len(pub.subscribers))
	}

	// Remove subscriber
	go func() { pub.removeSubCh <- sub }()
	select {
	case <-closed:
		// Subscriber was closed, this is expected
	case <-time.After(1 * time.Second):
		t.Error("Timed out waiting for subscriber closure")
	}

	if len(pub.subscribers) != 0 {
		t.Errorf("Expected 0 subscribers but got %d", len(pub.subscribers))
	}

	// Stop the publisher
	pub.Stop()
}
