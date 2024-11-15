package twosumarraysorted

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Test 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			name:     "Test 2",
			nums:     []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			name:     "Test 3",
			nums:     []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
