package delnode

import (
	"reflect"
	"testing"
)

// TestDeleteNode is the test function for DeleteNode using table-driven tests.
func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		nodeVal  int
		expected []int
	}{
		{
			name:     "Test 1",
			input:    []int{4, 5, 1, 9},
			nodeVal:  5,
			expected: []int{4, 1, 9},
		},
		{
			name:     "Test 2",
			input:    []int{4, 5, 1, 9},
			nodeVal:  1,
			expected: []int{4, 5, 9},
		},
		{
			name:     "Test 3",
			input:    []int{1, 2, 3, 4},
			nodeVal:  3,
			expected: []int{1, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			// Find the node to be deleted
			var nodeToDelete *ListNode
			for current := head; current != nil; current = current.Next {
				if current.Val == tt.nodeVal {
					nodeToDelete = current
					break
				}
			}
			// Delete the node
			DeleteNode(nodeToDelete)
			// Convert the result list back to a slice
			result := listToSlice(head)
			// Compare the result with the expected output
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DeleteNode() = %v, want %v", result, tt.expected)
			}
		})
	}
}
