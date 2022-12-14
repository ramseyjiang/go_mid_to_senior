package mergechans

import (
	"testing"
)

func TestMergeChans(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	aSlice := []int{1, 2, 3}
	bSlice := []int{20, 10, 30}
	cSlice := []int{300, 200, 100}

	go func() {
		for _, num := range aSlice {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range bSlice {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range cSlice {
			c <- num
		}
		close(c)
	}()

	tempWant := append(aSlice, bSlice...)
	want := append(tempWant, cSlice...)
	var expected []int
	for num := range joinChannels(a, b, c) {
		expected = append(expected, num)
	}

	if len(want) != len(expected) {
		t.Error("Length wrong")
	}

	for _, v := range want {
		if !contains(expected, v) {
			t.Error("content wrong")
		}
	}
}
