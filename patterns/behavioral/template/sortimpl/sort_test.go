package sortimpl

import (
	"fmt"
	"sort"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSort(t *testing.T) {
	list := baseAlgorithm{
		intSlice: []int{5, 8, 2, 1},
	}
	expected := baseAlgorithm{
		intSlice: []int{1, 2, 5, 8},
	}
	sort.Sort(list)
	fmt.Println(list)

	assert.Equal(t, list, expected)
}
