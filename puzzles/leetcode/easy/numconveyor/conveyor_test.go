package numconveyor

import "testing"

func TestConveyor(t *testing.T) {
	expected := []int{0, 1, 4, 9, 16}
	want := conveyor()

	if len(expected) != len(want) {
		t.Error("Length wrong")
	}

	for _, v := range want {
		if !contains(expected, v) {
			t.Error("content wrong")
		}
	}
}
