package containermostwater

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			"Test 1",
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			"Test 2",
			[]int{1, 1},
			1,
		},
		{
			"Test 3",
			[]int{4, 3, 2, 1, 4},
			16,
		},
		{
			"Test 4",
			[]int{1, 2, 1},
			2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maxArea(test.height)
			if result != test.expected {
				t.Errorf("For height %v, expected %d, but got %d", test.height, test.expected, result)
			}
		})
	}
}
