package superior

import (
	"testing"
)

func TestTableDriven(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, tc := range testCases {
		result := add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestSubtests(t *testing.T) {
	t.Run("Positive numbers", func(t *testing.T) {
		result := add(1, 2)
		expected := 3
		if result != expected {
			t.Errorf("add(1, 2) = %d; want %d", result, expected)
		}
	})

	t.Run("Zeros", func(t *testing.T) {
		result := add(0, 0)
		expected := 0
		if result != expected {
			t.Errorf("add(0, 0) = %d; want %d", result, expected)
		}
	})
}

func TestCombine(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positive numbers", 1, 2, 3},
		{"Zeros", 0, 0, 0},
		{"Negative and positive numbers", -1, 1, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
