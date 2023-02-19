package publisher

import "fmt"

type Observer interface {
	Notify(string)
}

// Publisher struct, which is the one that triggers an event, so it must accept new observers and remove them if necessary.
// When the Publisher struct is triggered, it must notify all its observers of the new event with the data associated.
// Publisher structure stores the list of subscribed observers in a slice field called ObserversList.
type Publisher struct {
	ObserversList []Observer
}

func (s *Publisher) AddObserver(o Observer) {
	s.ObserversList = append(s.ObserversList, o)
}

func (s *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int

	for i, observer := range s.ObserversList {
		// comparing the Observer object's o variable with the ones stored in the list
		if observer == o {
			indexToRemove = i
			break
		}
	}

	// 1. First, we need to use slice indexing to return a new slice containing every object from the beginning of the slice to the index we want to remove (not included).
	// 2. Then, we get another slice from the index we want to remove (not included) to the last object in the slice
	// 3. Finally, we join the previous two new slices into a new one (the append function)
	// For example, in a list from 1 to 10 in which we want to remove the number 5,
	// we have to create a new slice, joining a slice from 1 to 4 and a slice from 6 to 10.
	s.ObserversList = append(s.ObserversList[:indexToRemove], s.ObserversList[indexToRemove+1:]...)
}

// NotifyObservers method with a string that acts as the message we want to spread between all observers.
func (s *Publisher) NotifyObservers(m string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", m)
	for _, observer := range s.ObserversList {
		observer.Notify(m)
	}
}
