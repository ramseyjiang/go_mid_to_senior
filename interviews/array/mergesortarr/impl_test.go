package mergesortarr

import (
	"reflect"
	"testing"
)

func TestMergeSortArr(t *testing.T) {
	testCases := []struct {
		name      string
		inputArr1 []int
		inputNum1 int
		inputArr2 []int
		inputNum2 int
		expected  []int
	}{
		{
			"test 1",
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			"test 2",
			[]int{1},
			1,
			[]int{0},
			0,
			[]int{1},
		},
		{
			"test 3",
			[]int{0},
			0,
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mergeSortedArr(tt.inputArr1, tt.inputNum1, tt.inputArr2, tt.inputNum2)
			if !reflect.DeepEqual(tt.inputArr1, tt.expected) {
				t.Errorf("got %v, want %v", tt.inputArr1, tt.expected)
			}
		})
	}
}
