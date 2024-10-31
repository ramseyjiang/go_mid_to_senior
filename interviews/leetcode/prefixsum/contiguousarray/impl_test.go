package contiguousarray

import (
	"testing"
)

func TestFindMaxLength(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Example 1", []int{0, 1}, 2},
		{"Example 2", []int{0, 1, 0}, 2},
		{"No equal subarray", []int{1, 1, 1}, 0},
		{"All zeros", []int{0, 0, 0}, 0},
		{"Mixed equal parts", []int{0, 1, 1, 0, 1, 0}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMaxLength(tt.input)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
