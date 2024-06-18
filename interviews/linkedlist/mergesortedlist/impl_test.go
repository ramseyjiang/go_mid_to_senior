package mergesortedlist

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	list1    []int
	list2    []int
	expected []int
}{
	{
		name:     "Example 1",
		list1:    []int{1, 2, 4},
		list2:    []int{1, 3, 4},
		expected: []int{1, 1, 2, 3, 4, 4},
	},
	{
		name:     "Example 2",
		list1:    []int{},
		list2:    []int{},
		expected: []int{},
	},
	{
		name:     "Example 3",
		list1:    []int{},
		list2:    []int{0},
		expected: []int{0},
	},
}

func TestMergeListsIterative(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := sliceToList(tt.list1)
			list2 := sliceToList(tt.list2)
			mergedHead := mergeListsIterative(list1, list2)
			result := listToSlice(mergedHead)

			if len(result) == 0 && len(tt.expected) == 0 {
				// Both are empty, test passes
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mergeTwoLists() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMergeListsRecursive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := sliceToList(tt.list1)
			list2 := sliceToList(tt.list2)
			mergedHead := mergeListsRecursive(list1, list2)
			result := listToSlice(mergedHead)

			if len(result) == 0 && len(tt.expected) == 0 {
				// Both are empty, test passes
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mergeTwoLists() = %v, want %v", result, tt.expected)
			}
		})
	}
}
