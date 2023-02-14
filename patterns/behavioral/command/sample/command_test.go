package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// The client invokes a command using an Invoker, which calls the command's execute method,
// which, in turn, delegates to the Receiver's action method.
func TestExecuteCommand(t *testing.T) {
	receiver := &ConcreteReceiver{}
	command := &ConcreteCommand{receiver: receiver}
	invoker := &Invoker{command: command}
	assert.Equal(t, "Action called", invoker.ExecuteCommand())
}
