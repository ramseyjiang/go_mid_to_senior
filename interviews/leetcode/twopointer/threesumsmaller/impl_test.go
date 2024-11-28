package threeSumSmaller

import (
	"testing"
)

func TestThreeSumSmaller(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			"test 1",
			[]int{-2, 0, 1, 3},
			2,
			2,
		}, // Two valid triplets: [-2, 0, 1] and [-2, 0, 3]
		{
			"test 2",
			[]int{},
			0,
			0,
		}, // No elements, so no triplets
		{
			"test 3",
			[]int{0},
			0,
			0,
		}, // Single element, no triplets
		{
			"test 4",
			[]int{3, 1, -2, 0},
			4,
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumSmaller(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("threeSumSmaller(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
