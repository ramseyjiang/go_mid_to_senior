package rangesumquery

import (
	"testing"
)

func TestNumArray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		query    []int
		expected int
	}{
		{
			name:     "Test 1",
			nums:     []int{-2, 0, 3, -5, 2, -1},
			query:    []int{0, 2},
			expected: 1,
		},
		{
			name:     "Test 2",
			nums:     []int{-2, 0, 3, -5, 2, -1},
			query:    []int{2, 5},
			expected: -1,
		},
		{
			name:     "Test 3",
			nums:     []int{-2, 0, 3, -5, 2, -1},
			query:    []int{0, 5},
			expected: -3,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			prefixSum := PrefixSum(tt.nums)
			got := prefixSum.SumRange(tt.query[0], tt.query[1])
			if got != tt.expected {
				t.Errorf("SumRange(%d, %d) = %d; expected %d", tt.query[0], tt.query[1], got, tt.expected)
			}
		})
	}
}
