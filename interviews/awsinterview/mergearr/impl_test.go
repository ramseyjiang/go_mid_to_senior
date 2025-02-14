package mergearr

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "Example 1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Example 2",
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "No Overlap",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected:  [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:      "Nested Intervals",
			intervals: [][]int{{1, 10}, {2, 6}, {8, 9}},
			expected:  [][]int{{1, 10}},
		},
		{
			name:      "Single Interval",
			intervals: [][]int{{1, 5}},
			expected:  [][]int{{1, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := merge(tt.intervals)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
