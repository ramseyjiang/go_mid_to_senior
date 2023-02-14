package noreceiver

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAddCommand(t *testing.T) {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))
	assert.Equal(t, "First message", queue.record[0])
	assert.Equal(t, "Second message", queue.record[1])
	assert.Equal(t, "Third message", queue.record[2])

	// because queue defined length 3, the Fourth message and Fifth message cannot be inserted into.
	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))
	assert.NotEqual(t, "Fourth message", queue.record[0])
	assert.NotEqual(t, "Fifth message", queue.record[1])
}
