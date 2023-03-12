package log

import "fmt"

type Logger interface {
	Log(message string) string
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) string {
	return message
}

type Person struct {
	Logger
	Name string
}

func (p Person) SayHello() string {
	return p.Log(fmt.Sprintf("Hello, my name is %s", p.Name))
}
