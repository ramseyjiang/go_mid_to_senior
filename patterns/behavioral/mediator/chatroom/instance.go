package chatroom

import "fmt"

// The Mediator interface defines a method for sending messages to colleagues
type Mediator interface {
	SendMessage(message string, colleague Colleague) string
}

// Colleague interface defines methods for GetName and ReceiveMessage.
type Colleague interface {
	GetName() string
	ReceiveMessage(message string)
}

// The Chatroom is a concrete type, it implements the mediator interface.
// the Chatroom is a mediator between the users, receiving and directing messages between them.
type Chatroom struct {
	colleagues []Colleague
}

func (c *Chatroom) SendMessage(message string, sender Colleague) string {
	return sender.GetName() + ": " + message
}

// User is a concrete colleague implementation of the Colleague interface.
// The User objects only need to know about the Mediator interface and the Colleague interface, and not each other directly.
// This promotes loose coupling between the objects and simplifies their relationships.
type User struct {
	name     string
	mediator Mediator
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) ReceiveMessage(message string) {
	fmt.Printf("%s received: %s\n", u.name, message)
}

func (u *User) SendMessage(message string) string {
	return u.mediator.SendMessage(message, u)
}
