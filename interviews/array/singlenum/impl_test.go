package singlenum

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Example 1", []int{2, 2, 1}, 1},
		{"Example 2", []int{4, 1, 2, 1, 2}, 4},
		{"Example 3", []int{1}, 1},
		// You can add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.nums)
			if got != tt.want {
				t.Errorf("singleNumber(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
