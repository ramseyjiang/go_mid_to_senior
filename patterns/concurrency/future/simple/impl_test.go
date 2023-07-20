package simple

import "testing"

func TestCompute(t *testing.T) {
	t.Run("TestComputeWith5", func(t *testing.T) {
		future := compute(5)
		result := <-future
		if result != 10 {
			t.Errorf("Expected 10, but got %d", result)
		}
	})

	t.Run("TestComputeWith7", func(t *testing.T) {
		future := compute(7)
		result := <-future
		if result != 14 {
			t.Errorf("Expected 14, but got %d", result)
		}
	})
}
