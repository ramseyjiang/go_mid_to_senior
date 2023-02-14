package corc

import (
	"strings"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

// we create a variable called second with a pointer to a Logger; this is going to
// be the second link in our chain. Then we create a variable called first, that will be the first
// link in the chain. The first link points to the second variable, the second link in the chain.
func TestTimePassed(t *testing.T) {
	secondChain := new(Logger)
	firstChain := Logger{NextChain: secondChain}

	command := &TimePassed{start: time.Now()}

	firstChain.Next(command)

	assert.Equal(t, true, strings.Contains(firstChain.record[0], "Elapsed time from creation: "))
	assert.Equal(t, true, strings.Contains(secondChain.record[0], "Elapsed time from creation: "))
}
