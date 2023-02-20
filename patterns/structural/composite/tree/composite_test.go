package tree

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
)

type OperationLogger struct {
	*bytes.Buffer
}

// create a Leaf and a Composite.
// We add the Leaf to the Composite using the Add() method, and then call the Operation() method on the Composite
func TestComposite(t *testing.T) {
	leaf1 := &Leaf{}
	leaf2 := &Leaf{}
	composite := &Composite{}

	t.Run("Test default composite", func(t *testing.T) {
		assert.Equal(t, 0, len(composite.children))
		if len(composite.children) != 0 {
			t.Errorf("Expected 0 children, got %d", len(composite.children))
		}
	})

	// Test add method
	t.Run("Test add component", func(t *testing.T) {
		composite.Add(leaf1)
		composite.Add(leaf2)
		assert.Equal(t, 2, len(composite.children))
		if len(composite.children) != 2 {
			t.Errorf("Expected 2 children, got %d", len(composite.children))
		}
	})

	// Because t.Run() is a test closure.
	// When you run the test closure only, if you don't execute add() first, the default composite does not have any components.
	// If you run the whole TestComposite, the result always different with run the test closure only.
	t.Run("Test remove component", func(t *testing.T) {
		composite.Add(leaf1)
		composite.Add(leaf2)
		composite.Remove(leaf2)

		assert.Equal(t, 3, len(composite.children))
		if len(composite.children) != 3 {
			t.Errorf("Expected 1 child, got %d", len(composite.children))
		}
	})

	t.Run("Test operation method", func(t *testing.T) {
		leaf3 := &Leaf{}
		composite.Add(leaf3)

		logger := &OperationLogger{Buffer: &bytes.Buffer{}}
		log.SetOutput(logger)

		composite.Operation()

		// Check the output length, it is easy for doing the test.
		// Because the log output always includes the current time, it is hard to test for the content.
		if len(logger.String()) != 180 {
			t.Errorf("Unexpected output: %s", logger.String())

			fmt.Printf("Actual bytes: %v\n", logger.String())
		}
	})
}
