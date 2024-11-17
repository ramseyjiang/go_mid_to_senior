package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     [][]int
		expected   []interface{}
	}{
		{
			name:       "Test 1: Basic operations",
			operations: []string{"TwoSum", "Add", "Add", "Add", "Find", "Find"},
			values:     [][]int{{}, {1}, {3}, {5}, {4}, {7}},
			expected:   []interface{}{nil, nil, nil, nil, true, false},
		},
		{
			name:       "Test 2: Duplicates handling",
			operations: []string{"TwoSum", "Add", "Add", "Add", "Add", "Find"},
			values:     [][]int{{}, {1}, {1}, {3}, {3}, {6}},
			expected:   []interface{}{nil, nil, nil, nil, nil, true},
		},
		{
			name:       "Test 3: Pair of zeros",
			operations: []string{"TwoSum", "Add", "Add", "Find"},
			values:     [][]int{{}, {0}, {0}, {0}},
			expected:   []interface{}{nil, nil, nil, true},
		},
		{
			name:       "Test 4: Negative numbers",
			operations: []string{"TwoSum", "Add", "Add", "Add", "Find", "Find"},
			values:     [][]int{{}, {-1}, {-3}, {4}, {1}, {-2}},
			expected:   []interface{}{nil, nil, nil, nil, true, false},
		},
		{
			name:       "Test 5: No pair found",
			operations: []string{"TwoSum", "Add", "Add", "Find"},
			values:     [][]int{{}, {10}, {20}, {5}},
			expected:   []interface{}{nil, nil, nil, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var twoSum TwoSum
			for i, operation := range tt.operations {
				switch operation {
				case "TwoSum":
					twoSum = Constructor() // Initialize a new instance
				case "Add":
					twoSum.Add(tt.values[i][0])
				case "Find":
					result := twoSum.Find(tt.values[i][0])
					expectedResult := tt.expected[i].(bool)
					if result != expectedResult {
						t.Errorf("Find(%v) = %v; expected %v", tt.values[i][0], result, expectedResult)
					}
				}
			}
		})
	}
}
