package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestExecuteCommand(t *testing.T) {
	receiver := &ConcreteReceiver{}
	command := &ConcreteCommand{receiver: receiver}
	invoker := &Invoker{command: command}
	assert.Equal(t, "Action called", invoker.ExecuteCommand())
}
