package findduplicate

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			name:     "Example 3",
			nums:     []int{3, 3, 3, 3},
			expected: 3,
		},
		{
			name:     "Edge Case 1",
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			name:     "Edge Case 2",
			nums:     []int{1, 2, 2},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
