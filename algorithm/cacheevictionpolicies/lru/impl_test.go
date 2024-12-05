package lru

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	testCases := []struct {
		name       string
		operations []string
		values     [][]int
		expected   []int
	}{
		{
			name:       "Basic operations with eviction",
			operations: []string{"Put", "Put", "Get", "Put", "Get", "Get", "Get", "Get"},
			values:     [][]int{{1, 5}, {2, 2}, {1}, {3, 3}, {2}, {1}, {3}, {4}},
			expected:   []int{-1, -1, 5, -1, -1, 5, 3, -1}, // -1 for "Put" results
		},
		{
			name:       "Edge case with single capacity",
			operations: []string{"Put", "Get", "Put", "Get", "Get"},
			values:     [][]int{{1, 10}, {1}, {2, 20}, {1}, {2}},
			expected:   []int{-1, 10, -1, 10, 20}, // -1 for "Put" results
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewCacheLRU(2)
			var results []int
			for i, op := range tt.operations {
				switch op {
				case "Put":
					cache.Put(tt.values[i][0], tt.values[i][1])
					results = append(results, -1) // No return value for "Put"
				case "Get":
					results = append(results, cache.Get(tt.values[i][0]))
				}
			}
			fmt.Println(results)
			// Validate results
			for i, exp := range tt.expected {
				if results[i] != exp {
					t.Errorf("operation %v with value %v: expected %d, got %d", tt.operations[i], tt.values[i], exp, results[i])
				}
			}
		})
	}
}
