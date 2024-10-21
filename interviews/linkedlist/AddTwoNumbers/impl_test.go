package AddTwoNumbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input1   []int
		input2   []int
		expected []int
	}{
		{
			"test 1",
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 8},
		},
		{"test 2",
			[]int{0},
			[]int{0},
			[]int{0},
		},
		{"test 3",
			[]int{9, 9, 9, 9, 9, 9, 9},
			[]int{9, 9, 9, 9},
			[]int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := sliceToList(tt.input1)
			l2 := sliceToList(tt.input2)
			result := addTwoNumbers(l1, l2)
			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("got %v, expected %v", resultSlice, tt.expected)
			}
		})
	}
}
