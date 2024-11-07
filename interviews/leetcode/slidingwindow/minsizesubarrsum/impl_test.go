package minsizesubarrsum

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{"Example 1", 7, []int{2, 3, 1, 2, 4, 3}, 2},
		{"Example 2", 4, []int{1, 4, 4}, 1},
		{"Example 3", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{"Larger sum", 15, []int{1, 2, 3, 4, 5}, 5},
		{"Exact match", 11, []int{1, 2, 3, 4, 5}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubArrLen(tt.target, tt.nums)
			if got != tt.want {
				t.Errorf("minSubArrLen(%d, %v) = %d; want %d", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}
