package flipimage

import (
	"reflect"
	"testing"
)

// TestFlipAndInvertImage tests the flipAndInvertImage function.
func TestFlipAndInvertImage(t *testing.T) {
	testCases := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		{
			name:   "Example 1",
			input:  [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			expect: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			name:   "Example 2",
			input:  [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			expect: [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
		{
			name:   "Single row",
			input:  [][]int{{1, 0, 1}},
			expect: [][]int{{0, 1, 0}},
		},
		{
			name:   "Single column",
			input:  [][]int{{1}, {0}, {1}},
			expect: [][]int{{0}, {1}, {0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := flipInvertImage(tc.input)
			if !reflect.DeepEqual(result, tc.expect) {
				t.Errorf("flipInvertImage(%v) = %v, want %v", tc.input, result, tc.expect)
			}
		})
	}
}
