package tracffic

import (
	"testing"
)

func TestStateTransitions(t *testing.T) {
	tests := []struct {
		name            string
		startState      State
		expectedMessage string
		expectedNext    string
	}{
		{
			name:            "RedLightState",
			startState:      new(RedLightState),
			expectedMessage: "Red Light - Stop!",
			expectedNext:    "Yellow Light - Wait!",
		},
		{
			name:            "YellowLightState",
			startState:      new(YellowLightState),
			expectedMessage: "Yellow Light - Wait!",
			expectedNext:    "Green Light - Go!",
		},
		{
			name:            "GreenLightState",
			startState:      new(GreenLightState),
			expectedMessage: "Green Light - Go!",
			expectedNext:    "Red Light - Stop!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			context := NewContext(test.startState)
			message := context.Request()
			if message != test.expectedMessage {
				t.Errorf("State message was incorrect, got: %s, want: %s.", message, test.expectedMessage)
			}
			nextMessage := context.Request()
			if nextMessage != test.expectedNext {
				t.Errorf("Next state message was incorrect, got: %s, want: %s.", nextMessage, test.expectedNext)
			}
		})
	}
}
