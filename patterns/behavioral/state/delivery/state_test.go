package delivery

import (
	"testing"
)

func TestPackage(t *testing.T) {
	tests := []struct {
		name     string
		state    State
		expected string
	}{
		{
			name:     "Test Ordered State",
			state:    &OrderedState{},
			expected: "Package is ordered and is waiting for courier.",
		},
		{
			name:     "Test Shipped State",
			state:    &ShippedState{},
			expected: "Package is shipped and is in transit.",
		},
		{
			name:     "Test Delivered State",
			state:    &DeliveredState{},
			expected: "Package is delivered to the recipient.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPackage()
			p.UpdateState(tt.state)
			result := p.CurrentState()

			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}
