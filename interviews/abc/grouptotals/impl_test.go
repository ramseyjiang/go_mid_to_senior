package grouptotals

import (
	"testing"
)

func TestGroupTotals(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			"Basic test",
			[]string{"B:-1", "A:1", "B:3", "A:5"},
			"A:6,B:2",
		},
		{
			"Exclude zero test",
			[]string{"Z:0", "A:-1"},
			"A:-1",
		},
		{
			"Multiple test",
			[]string{"B:-1", "C:6", "B:3", "A:5", "D:-5", "D:5"},
			"A:5,B:2,C:6",
		},
		{
			"All zero test",
			[]string{"Z:0", "A:0", "D:-5", "D:5"},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GroupTotals(tt.input)
			if result != tt.expected {
				t.Errorf("GroupTotals() = %v, want %v", result, tt.expected)
			}
		})
	}
}
