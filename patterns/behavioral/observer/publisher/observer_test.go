package publisher

import (
	"testing"
)

func TestPublisher(t *testing.T) {
	tests := []struct {
		name         string
		subscribers  []Observer
		message      string
		expectedLogs map[string][]string
	}{
		{
			name: "Single Subscriber",
			subscribers: []Observer{
				&ConcreteObserver{Name: "Observer 1"},
			},
			message: "Hello, World!",
			expectedLogs: map[string][]string{
				"Observer 1": {"Hello, World!"},
			},
		},
		{
			name: "Multiple Subscribers",
			subscribers: []Observer{
				&ConcreteObserver{Name: "Observer 1"},
				&ConcreteObserver{Name: "Observer 2"},
			},
			message: "Hello, World!",
			expectedLogs: map[string][]string{
				"Observer 1": {"Hello, World!"},
				"Observer 2": {"Hello, World!"},
			},
		},
		{
			name:         "No Subscribers",
			subscribers:  []Observer{},
			message:      "Hello, World!",
			expectedLogs: map[string][]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			publisher := &ConcreteSubject{}

			for _, observer := range test.subscribers {
				publisher.RegisterObserver(observer)
			}

			publisher.NotifyObservers(test.message)

			for _, observer := range test.subscribers {
				expectedLogs := test.expectedLogs[observer.(*ConcreteObserver).Name]
				receivedLogs := observer.(*ConcreteObserver).Messages

				if len(receivedLogs) != len(expectedLogs) {
					t.Errorf("expected %d log(s) for observer %s, got %d", len(expectedLogs), observer.(*ConcreteObserver).Name, len(receivedLogs))
				}

				for i := range receivedLogs {
					if receivedLogs[i] != expectedLogs[i] {
						t.Errorf("observer %s: expected log %q, got %q", observer.(*ConcreteObserver).Name, expectedLogs[i], receivedLogs[i])
					}
				}
			}
		})
	}
}
