package rotate

import (
	"reflect"
	"testing"
)

// TestRotate is the unit test for the rotate function.
func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{"rotate by 3", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"rotate by 2", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{"rotate by 0", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"rotate by length", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"rotate by more than length", []int{1, 2, 3}, 5, []int{2, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.expected) {
				t.Errorf("rotate(%v, %d) got %v, want %v", tt.nums, tt.k, tt.nums, tt.expected)
			}
		})
	}
}
