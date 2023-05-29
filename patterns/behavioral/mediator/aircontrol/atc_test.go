package aircontrol

import (
	"testing"
)

func TestATC(t *testing.T) {
	atc := &ATCImpl{}

	airplane1 := &AirplaneImpl{name: "Airplane1", atc: atc}
	airplane2 := &AirplaneImpl{name: "Airplane2", atc: atc}

	atc.Register(airplane1)
	atc.Register(airplane2)

	tests := []struct {
		name     string
		sender   Airplane
		receiver *AirplaneImpl
		message  string
	}{
		{"Airplane1 sends to Airplane2", airplane1, airplane2, "Requesting to land."},
		{"Airplane2 sends to Airplane1", airplane2, airplane1, "Requesting to take off."},
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
