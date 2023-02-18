package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestState(t *testing.T) {
	content := &Content{state: &StateA{}}
	assert.Equal(t, "State A", content.Request())

	content.SetState(&StateB{})
	assert.Equal(t, "State B", content.Request())
}
