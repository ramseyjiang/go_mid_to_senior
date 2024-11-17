package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Test 1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "Test 2",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "Test 3",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "Test 4",
			nums:     []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSum(tt.nums)
			if !isEqual(result, tt.expected) {
				t.Errorf("threeSum(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
			// if !reflect.DeepEqual(result, tt.expected) {
			// 	t.Errorf("threeSum(%v) = %v, expected %v", tt.nums, result, tt.expected)
			// }
		})
	}
}

func isEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) || !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}
