package nearestsmallernum

import (
	"reflect"
	"testing"
)

func TestNearestSmallerValues(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"Basic case",
			[]int{5, 3, 1, 9, 7},
			[]int{-1, -1, -1, 1, 1},
		},
		{
			"Descending order",
			[]int{5, 4, 3, 2, 1},
			[]int{-1, -1, -1, -1, -1},
		},
		{
			"Ascending order",
			[]int{2, 4, 5, 1, 7},
			[]int{-1, 2, 4, -1, 1},
		},
		{
			"Mixed numbers",
			[]int{5, 3, 1, 9, 7},
			[]int{-1, -1, -1, 1, 1},
		},
		{
			"Single num",
			[]int{10},
			[]int{-1},
		},
		{
			"Empty array",
			[]int{},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := NearestSmallerValues(tt.input)

			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("output: %v, expected: %v", output, tt.expected)
			}
		})
	}
}
