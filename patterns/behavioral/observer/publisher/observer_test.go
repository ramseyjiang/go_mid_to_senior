package publisher

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

type TestObserver struct {
	ID      int
	Message Event
}

// TestObserver structure implements the Observer pattern by defining Notify(string) method
func (p *TestObserver) Notify(e Event) {
	p.Message = e
}

// given a different ID to each observer so that we can see later that each of them has printed the expected message.
func TestSubject(t *testing.T) {
	testObserver1 := &TestObserver{1, Event{"default"}}
	testObserver2 := &TestObserver{2, Event{"default"}}
	testObserver3 := &TestObserver{3, Event{"default"}}

	publisher := Publisher{}

	// add the observers by calling the AddObserver method on the Publisher structure.
	t.Run("AddObserver", func(t *testing.T) {
		publisher.AddObserver(testObserver1)
		publisher.AddObserver(testObserver2)
		publisher.AddObserver(testObserver3)

		assert.Equal(t, len(publisher.ObserversList), 3)
		if len(publisher.ObserversList) != 3 {
			t.Fail()
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		// RemoveObserver takes the observer with ID 2 and remove it from the list.
		publisher.RemoveObserver(testObserver2)

		assert.Equal(t, len(publisher.ObserversList), 2)

		for _, observer := range publisher.ObserversList {
			// make a type casting from a pointer to an observer, to a pointer to the TestObserver structure,
			// and check that the casting has been done correctly.
			testObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail() // This marks the test as failed, but allows the execution to continue.
			}

			assert.NotEqual(t, testObserver.ID, 2)
			// check that none of the observers left have the ID 2 because it must be removed.
			if testObserver.ID == 2 {
				t.Fail()
			}
		}
	})

	// When using the Notify method, all instances of TestObserver structure must change their Message field from empty to the passed message
	t.Run("Notify", func(t *testing.T) {
		assert.NotEqual(t, len(publisher.ObserversList), 0)

		for _, observer := range publisher.ObserversList {
			// make a type casting from a pointer to an observer, to a pointer to the TestObserver structure,
			// and check that the casting has been done correctly.
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			assert.Equal(t, printObserver.Message, Event{"default"})
		}

		message := Event{"Hello World!"}
		publisher.NotifyObservers(Event{"Hello World!"})

		for _, observer := range publisher.ObserversList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			assert.Equal(t, printObserver.Message, message)
		}
	})
}
