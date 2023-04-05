package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// The client invokes a command using an Invoker, which calls the command's execute method,
// which, in turn, delegates to the Receiver's action method.
func TestExecuteCommand(t *testing.T) {
	receiver := &Receiver{}
	commandA := &ConcreteCommandA{receiver: receiver}
	commandB := &ConcreteCommandB{receiver: receiver}

	invoker := &Invoker{}
	invoker.SetCommand(commandA)
	assert.Equal(t, "ActionA called", invoker.ExecuteCommand())

	invoker.SetCommand(commandB)
	assert.Equal(t, "ActionB called", invoker.ExecuteCommand())
}
