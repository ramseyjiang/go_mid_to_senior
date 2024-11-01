package subarrequalk

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"Example 1", []int{1, 1, 1}, 2, 2},
		{"Example 2", []int{1, 2, 3}, 3, 2},
		{"No Subarrays with Sum k", []int{1, 2, 3}, 7, 0},
		{"Multiple Subarrays with Sum k", []int{1, -1, 1, 1, -1, 1}, 2, 4},
		{"All Zeros with k Zero", []int{0, 0, 0, 0}, 0, 10}, // Every subarray here sums to 0
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubarraySum(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
