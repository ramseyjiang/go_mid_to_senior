package aircontrol

import "fmt"

// ATC is the mediator interface
type ATC interface {
	Register(airplane Airplane)
	Send(message string, origin Airplane)
}

// Airplane is the colleague interface
type Airplane interface {
	Send(message string)
	Receive(message string, origin Airplane)
	GetName() string
	GetMessages() []string
}

// ATCImpl is the concrete mediator
type ATCImpl struct {
	airplanes []Airplane
}

func (a *ATCImpl) Register(airplane Airplane) {
	a.airplanes = append(a.airplanes, airplane)
}

func (a *ATCImpl) Send(message string, origin Airplane) {
	for _, ap := range a.airplanes {
		if ap != origin {
			ap.Receive(message, origin)
		}
	}
}

// AirplaneImpl is the concrete colleague
type AirplaneImpl struct {
	name     string
	atc      ATC
	messages []string
}

func (ap *AirplaneImpl) Send(message string) {
	fmt.Println(ap.name + " sends message: " + message)
	ap.atc.Send(message, ap)
}

func (ap *AirplaneImpl) Receive(message string, origin Airplane) {
	fmt.Println(ap.name + " received message from " + origin.GetName() + ": " + message)
	ap.messages = append(ap.messages, message)
}

func (ap *AirplaneImpl) GetName() string {
	return ap.name
}

func (ap *AirplaneImpl) GetMessages() []string {
	return ap.messages
}
