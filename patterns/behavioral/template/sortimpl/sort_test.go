package sortimpl

import (
	"reflect"
	"sort"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSort(t *testing.T) {
	var list intList = []int{5, 8, 2, 1}

	// here if defines expected := []int{1, 2, 5, 8}, it won't equal to list, because the type are different.
	// Hence, the only way to make reflect.DeepEqual is true, you should use MyList as a type.
	var expected intList = []int{1, 2, 5, 8}
	sort.Sort(list)

	assert.Equal(t, true, reflect.DeepEqual(list, expected))
}
