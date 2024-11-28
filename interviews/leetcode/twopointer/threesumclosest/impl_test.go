package threesumclosest

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"Test 1", []int{-1, 2, 1, -4}, 1, 2},
		{"Test 2", []int{0, 0, 0}, 1, 0},
		{"Test 3", []int{-1, 0, 1, 1}, 2, 2},
		{"Test 4", []int{-3, -2, -5, 3, -4}, -1, -2},
		{"Test 5", []int{1, 1, 1, 0}, -100, 2},
		{"Test 6", []int{1, 2, 4, 8, 16, 32, 64, 128}, 82, 82},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSumClosest(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("threeSumClosest(%v, %d) = %d; expected %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
