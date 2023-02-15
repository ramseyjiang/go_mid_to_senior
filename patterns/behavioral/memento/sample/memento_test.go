package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMemento(t *testing.T) {
	originator := &Originator{state: "Initial state"}
	assert.Equal(t, "Initial state", originator.getState())

	caretaker := &CareTaker{}
	caretaker.addMemento(originator.createMemento())

	originator.setState("First state")
	assert.Equal(t, "First state", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("Second state")
	assert.Equal(t, "Second state", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	assert.Equal(t, "First state", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	assert.Equal(t, "Initial state", originator.getState())
}
