package log

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLog(t *testing.T) {
	logger := ConsoleLogger{}
	person := Person{logger, "John"}

	assert.Equal(t, "Hello, my name is John", person.SayHello())
}
