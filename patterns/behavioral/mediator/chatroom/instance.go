package chatroom

import "fmt"

// Chatroom is the mediator interface
type Chatroom interface {
	Register(participant Participant)
	Send(message string, origin Participant)
}

// Participant is the colleague interface
type Participant interface {
	Send(message string)
	Receive(message string, origin Participant)
	GetName() string
}

// ChatroomImpl is the concrete mediator
type ChatroomImpl struct {
	participants []Participant
}

func (c *ChatroomImpl) Register(participant Participant) {
	c.participants = append(c.participants, participant)
}

func (c *ChatroomImpl) Send(message string, origin Participant) {
	for _, p := range c.participants {
		if p != origin {
			p.Receive(message, origin)
		}
	}
}

// ParticipantImpl is the concrete colleague
type ParticipantImpl struct {
	name     string
	chatroom Chatroom
	messages []string
}

func (p *ParticipantImpl) Send(message string) {
	fmt.Println(p.name + " sends message: " + message)
	p.chatroom.Send(message, p)
}

func (p *ParticipantImpl) Receive(message string, origin Participant) {
	fmt.Println(p.name + " received message from " + origin.GetName() + ": " + message)
	p.messages = append(p.messages, message)
}

func (p *ParticipantImpl) GetName() string {
	return p.name
}

func (p *ParticipantImpl) GetMessages() []string {
	return p.messages
}
