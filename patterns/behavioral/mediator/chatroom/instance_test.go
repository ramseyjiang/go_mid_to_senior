package chatroom

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMessage(t *testing.T) {
	chatroom := &Chatroom{}

	user1 := &User{name: "User 1", mediator: chatroom}
	user2 := &User{name: "User 2", mediator: chatroom}
	user3 := &User{name: "User 3", mediator: chatroom}

	assert.Equal(t, user1.name, user1.GetName())
	assert.Equal(t, user2.name, user2.GetName())
	assert.Equal(t, user3.name, user3.GetName())

	assert.Equal(t, "User 1: Hello, User 2!", user1.SendMessage("Hello, User 2!"))
	assert.Equal(t, "User 2: Hi, User 1!", user2.SendMessage("Hi, User 1!"))
	assert.Equal(t, "User 3: Hi, everyone!", user3.SendMessage("Hi, everyone!"))
}
