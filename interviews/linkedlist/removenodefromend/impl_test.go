package removenodefromend

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		head     []int
		n        int
		expected []int
	}{
		{
			name:     "Test 1",
			head:     []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},

		{
			name:     "Test 2",
			head:     []int{1},
			n:        1,
			expected: []int{},
		},
		{
			name:     "Test 3",
			head:     []int{1, 2},
			n:        1,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.head)
			result := removeNthFromEnd(head, tt.n)
			resultArr := listToSlice(result)

			if len(resultArr) == 0 && len(tt.expected) == 0 {
				// Both are empty, test passes
				return
			}

			if !reflect.DeepEqual(resultArr, tt.expected) {
				t.Errorf("removeNthFromEnd(%v, %d) = %v; expected %v", tt.head, tt.n, resultArr, tt.expected)
			}
		})
	}
}
