package reverselist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		head     []int
		expected []int
	}{
		{
			name:     "Test 1",
			head:     []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Test 2",
			head:     []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Test 3",
			head:     []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.head)
			result := reverseList(head)
			resultArr := listToSlice(result)

			if len(resultArr) == 0 && len(tt.expected) == 0 {
				// Both are empty, test passes
				return
			}

			if !reflect.DeepEqual(resultArr, tt.expected) {
				t.Errorf("reverseList(%v) = %v; expected %v", tt.head, resultArr, tt.expected)
			}
		})
	}
}
