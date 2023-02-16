package sample

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIterator(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	var copySlice []int
	iterator := NewSliceIntIterator(slice)
	for {
		value, ok := iterator.Next()
		if !ok {
			break
		}
		copySlice = append(copySlice, value)
	}
	fmt.Println(copySlice)
	assert.Equal(t, true, reflect.DeepEqual(slice, copySlice))
}
