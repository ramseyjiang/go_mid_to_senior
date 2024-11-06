package maxaveragesubarr

import "testing"

func TestFindMaxAverageArr(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{"Example 1", []int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
		{"Example 2", []int{5}, 1, 5.00000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxAverageArr(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("got %f, want %f", result, tt.expected)
			}
		})
	}
}
