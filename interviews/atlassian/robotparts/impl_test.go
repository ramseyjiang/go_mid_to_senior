package robotparts

import (
	"reflect"
	"testing"
)

func TestRobotParts(t *testing.T) {
	tests := []struct {
		name          string
		requiredParts string
		expected      []string
	}{
		{
			name:          "required_parts_1",
			requiredParts: "sensors,case,speaker,wheels",
			expected:      []string{"Rosie", "Optimus"},
		},
		{
			name:          "required_parts_2",
			requiredParts: "sensors,case,speaker,wheels,claw",
			expected:      []string{"Rosie"},
		},
		{
			name:          "required_parts_3",
			requiredParts: "sensors,case,screws",
			expected:      []string{},
		},
	}

	AllParts := []string{
		"Rosie_claw",
		"Rosie_sensors",
		"Dustie_case",
		"Optimus_sensors",
		"Rust_sensors",
		"Rosie_case",
		"Rust_case",
		"Optimus_speaker",
		"Rosie_wheels",
		"Rosie_speaker",
		"Dustie_case",
		"Dustie_arms",
		"Rust_claw",
		"Dustie_case",
		"Dustie_speaker",
		"Optimus_case",
		"Optimus_wheels",
		"Rust_legs",
		"Optimus_sensors",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := GetRobotParts(AllParts, tt.requiredParts)
			if !reflect.DeepEqual(output, tt.expected) {
				if len(output) == 0 && len(tt.expected) == 0 {
					// Both slices are empty (or nil), considered equal, no error
				} else {
					t.Errorf("YourFunction(%v) = %v; want %v", tt.requiredParts, output, tt.expected)
				}
			}
		})
	}
}
