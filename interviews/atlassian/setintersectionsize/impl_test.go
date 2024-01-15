package setintersectionsize

import "testing"

func TestIntersectionSize(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "Example 1",
			intervals: [][]int{{1, 3}, {3, 7}, {8, 9}},
			expected:  5,
		},
		{
			name:      "Example 2",
			intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
			expected:  3,
		},
		{
			name:      "Example 3",
			intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
			expected:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := intersectionSizeTwo(tt.intervals)
			if output != tt.expected {
				t.Errorf("intersectionSizeTwo() output is %v, expected is %v", output, tt.expected)
			}
		})
	}
}
