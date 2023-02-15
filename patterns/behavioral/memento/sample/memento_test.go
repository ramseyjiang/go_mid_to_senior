package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMemento(t *testing.T) {
	originator := &Originator{state: "Initial state"}
	assert.Equal(t, "Initial state", originator.getState())

	caretaker := &CareTaker{}
	currentLen := len(caretaker.mementos) // 0, because none mementos at the beginning
	caretaker.addMemento(originator.createMemento())
	assert.Equal(t, currentLen+1, len(caretaker.mementos))

	originator.setState("First state")
	assert.Equal(t, "First state", originator.getState())

	currentLen = len(caretaker.mementos)
	caretaker.addMemento(originator.createMemento())
	assert.Equal(t, currentLen+1, len(caretaker.mementos))

	originator.setState("Second state")
	assert.Equal(t, "Second state", originator.getState())

	currentLen = len(caretaker.mementos)
	caretaker.addMemento(originator.createMemento())
	assert.Equal(t, currentLen+1, len(caretaker.mementos))

	originator.restoreMemento(caretaker.getMemento(1))
	assert.Equal(t, "First state", originator.getState())
	assert.Equal(t, 3, len(caretaker.mementos))

	originator.restoreMemento(caretaker.getMemento(0))
	assert.Equal(t, "Initial state", originator.getState())
	assert.Equal(t, 3, len(caretaker.mementos))
}
