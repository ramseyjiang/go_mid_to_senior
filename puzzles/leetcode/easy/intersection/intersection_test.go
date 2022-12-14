package intersection

import (
	"testing"
)

func TestIntersection1(t *testing.T) {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}

	want := []int{2, 23}
	expected := intersection(a, b)

	if len(want) != len(expected) {
		t.Error("Wrong")
	}

	for i := range want {
		if want[i] != expected[i] {
			t.Error("Wrong")
		}
	}
}

func TestIntersection2(t *testing.T) {
	a := []int{1, 1, 1}
	b := []int{1, 1, 1, 1}

	want := []int{1, 1, 1}
	expected := intersection(a, b)

	if len(want) != len(expected) {
		t.Error("Wrong")
	}

	for i := range want {
		if want[i] != expected[i] {
			t.Error("Wrong")
		}
	}
}
