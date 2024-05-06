package majorityelement

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "test 1",
			input:  []int{1, 2, 3},
			output: -1,
		},
		{
			name:   "test 2",
			input:  []int{3, 1, 3, 3, 2},
			output: 3,
		},
		{
			name:   "test 3",
			input:  []int{3, 2, 3},
			output: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := majorityElement(tt.input)
			if expected != tt.output {
				t.Errorf("majorityElement(%v) = %v, want %v", tt.input, expected, tt.output)
			}
		})
	}
}
