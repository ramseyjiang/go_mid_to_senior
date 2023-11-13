package movezeros

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"Example 1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"Example 2", []int{0}, []int{0}},
		{"Example 3", []int{0, 1, 6, 3, 0, 12}, []int{1, 6, 3, 12, 0, 0}},
		// Additional test cases can be added here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("moveZeroes(%v) = %v, want %v", tt.nums, tt.nums, tt.want)
			}
		})
	}
}
