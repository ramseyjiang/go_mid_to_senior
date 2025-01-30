package strkeyboardlayout

import "testing"

func TestSpell(t *testing.T) {
	layout := [][]string{
		{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
		{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
		{"z", "x", "c", "v", "b", "n", "m"},
	}

	actual := spell(layout, "hello")
	expected := "DRRRRRPULLLPDRRRRRRPPUP"

	if actual != expected {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}
