package corc

import (
	"time"
)

// Command interface
type Command interface {
	Info() string
}

// TimePassed is a Concrete command
type TimePassed struct {
	start time.Time
}

// Info is the method implement the Command interface
func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

// ChainLogger is the Receiver interface. The ChainLogger is using chain of responsibility pattern.
type ChainLogger interface {
	Next(Command)
}

// Logger is a Concrete command, Logger structures implement the ChainLogger interface and represent the concrete handlers in the chain.
type Logger struct {
	NextChain ChainLogger
	record    []string
}

func (f *Logger) Next(c Command) {
	time.Sleep(time.Second)

	elapsedTime := "Elapsed time from creation: " + c.Info()
	f.record = append(f.record, elapsedTime)

	// checks whether it can handle the request and if not, passes the request to the next chain using the Next method.
	if f.NextChain != nil {
		f.NextChain.Next(c)
	}
}
