package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"Example 1", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Example 2", []int{3, 3}, 6, []int{1, 0}},
		{"Example 3", []int{2, 11, 7, 15}, 9, []int{0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !isValidResult(got, tt.nums, tt.target) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isValidResult(indices []int, nums []int, target int) bool {
	if len(indices) != 2 {
		return false
	}
	return nums[indices[0]]+nums[indices[1]] == target
}
