package simple

import (
	"testing"
	"time"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected float64
		hasError bool
	}{
		{
			name:     "Average of non-empty slice",
			data:     []int{10, 20, 30, 40, 50},
			expected: 30,
			hasError: false,
		},
		{
			name:     "Average of empty slice",
			data:     []int{},
			expected: 0,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			future := CalculateAverage(tt.data)

			// Using a select with a timeout to ensure the test doesn't hang indefinitely
			select {
			case <-time.After(3 * time.Second):
				t.Fatal("Test timed out")
			default:
				average, err := future.Get()
				if tt.hasError && err == nil {
					t.Fatalf("Expected an error but got none")
				}
				if !tt.hasError && err != nil {
					t.Fatalf("Didn't expect an error but got: %v", err)
				}
				if average != tt.expected {
					t.Fatalf("Expected average %f but got %f", tt.expected, average)
				}
			}
		})
	}
}
