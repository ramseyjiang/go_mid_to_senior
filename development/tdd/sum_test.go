package tdd

import (
	"fmt"
	"testing"
)

func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d+%d", tt.a, tt.b), func(t *testing.T) {
			result := Sum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Sum(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
