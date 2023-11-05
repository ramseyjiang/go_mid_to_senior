package bst

import (
	"bytes"
	"io"
	"testing"
)

// Helper function to capture output
func captureOutput(f func(w io.Writer)) string {
	var buf bytes.Buffer
	f(&buf)
	return buf.String()
}

func TestBST(t *testing.T) {
	// Initialize a new BST
	bst := BST{}

	// Test Insert
	t.Run("Insert", func(t *testing.T) {
		valuesToInsert := []int{10, 5, 15, 20, 17, 4, 6}
		for _, val := range valuesToInsert {
			bst.Insert(val)
		}

		// Check if root is correct
		if bst.root.data != 10 {
			t.Errorf("Expected root to be 10, got %d", bst.root.data)
		}
	})

	// Test Search
	t.Run("Search", func(t *testing.T) {
		testCases := []struct {
			desc  string
			value int
			want  bool
		}{
			{"Value exists", 5, true},
			{"Value does not exist", 11, false},
			{"Value is root", 10, true},
			{"Value is leaf", 4, true},
		}

		for _, tc := range testCases {
			t.Run(tc.desc, func(t *testing.T) {
				got := bst.Search(tc.value)
				if got != tc.want {
					t.Errorf("Search(%d) = %v; want %v", tc.value, got, tc.want)
				}
			})
		}
	})

	// Test InOrder
	t.Run("InOrder", func(t *testing.T) {
		expectedOutput := "4 5 6 10 15 17 20 "
		actualOutput := captureOutput(func(w io.Writer) {
			bst.InOrder(w, bst.root)
		})
		if actualOutput != expectedOutput {
			t.Errorf("Expected InOrder output to be %v, got %v", expectedOutput, actualOutput)
		}
	})

	// Test LevelOrder
	t.Run("LevelOrder", func(t *testing.T) {
		expectedOutput := "10 5 15 4 6 20 17 "
		actualOutput := captureOutput(func(w io.Writer) {
			bst.LevelOrder(w)
		})
		if actualOutput != expectedOutput {
			t.Errorf("Expected LevelOrder output to be %v, got %v", expectedOutput, actualOutput)
		}
	})
}
