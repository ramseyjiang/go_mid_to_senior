package publisher

// Step 1: Define a subject interface and an observer interface.

// Subject represents the subject being observed.
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(message string)
}

// Observer represents the observer.
type Observer interface {
	Update(message string)
}

// Step 2: Implement the subject interface.

// ConcreteSubject represents the concrete implementation of the Subject interface.
type ConcreteSubject struct {
	observers []Observer
}

// RegisterObserver adds an observer to the list of observers.
func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

// RemoveObserver removes an observer from the list of observers.
func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	indexToRemove := -1

	// Find the index of the observer to remove
	for i, obs := range s.observers {
		if obs == observer {
			indexToRemove = i
			break
		}
	}

	// Remove the observer from the list
	if indexToRemove >= 0 {
		s.observers = append(s.observers[:indexToRemove], s.observers[indexToRemove+1:]...)
	}
}

// NotifyObservers notifies all observers with the given message.
func (s *ConcreteSubject) NotifyObservers(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

// Step 3: Implement the observer interface.

// ConcreteObserver represents the concrete implementation of the Observer interface.
type ConcreteObserver struct {
	Name     string
	Messages []string
}

func (o *ConcreteObserver) Update(message string) {
	o.Messages = append(o.Messages, message)
}
