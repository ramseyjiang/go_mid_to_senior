package minwindowsubstr

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected string
	}{
		{"Test 1", "ADOBECODEBANC", "ABC", "BANC"},
		{"Test 2", "a", "a", "a"},
		{"Test 3", "a", "aa", ""},
		{"Test 4", "", "A", ""},
		{"Test 5", "acbbaca", "aba", "baca"},
		{"Test 6", "ab", "b", "b"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minSubstrWindow(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("MinWindow(%q, %q) = %q; expected %q", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}
