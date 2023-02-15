package sample

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	iterator := NewSliceIntIterator(slice)
	for {
		value, ok := iterator.Next()
		if !ok {
			break
		}
		fmt.Println(value)
	}
}
