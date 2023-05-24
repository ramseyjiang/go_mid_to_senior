package chatroom

import (
	"testing"
)

func TestChatroom(t *testing.T) {
	chatroom := &ChatroomImpl{}

	alice := &ParticipantImpl{name: "Alice", chatroom: chatroom}
	bob := &ParticipantImpl{name: "Bob", chatroom: chatroom}

	chatroom.Register(alice)
	chatroom.Register(bob)

	tests := []struct {
		name     string
		sender   Participant
		receiver *ParticipantImpl
		message  string
	}{
		{"Alice sends to Bob", alice, bob, "Hello, Bob!"},
		{"Bob sends to Alice", bob, alice, "Hello, Alice!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sender.Send(tt.message)
			messages := tt.receiver.GetMessages()
			if len(messages) == 0 || messages[len(messages)-1] != tt.message {
				t.Errorf("Expected to receive message %v, but it was not received", tt.message)
			}
		})
	}
}
