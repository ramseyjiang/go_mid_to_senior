package workerpoolfunc

import "testing"

func TestReuseRoutine(t *testing.T) {
	expected := []int{10, 20, 30, 40, 50, 60}
	want := reuseRoutine()

	if len(want) != len(expected) {
		t.Error("Length error.")
	}

	for _, v := range want {
		if !contains(expected, v) {
			t.Error("content wrong")
		}
	}
}
