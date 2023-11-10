package removeduplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "Example 2",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "All unique",
			input:    []int{1, 2, 3, 4},
			expected: 4,
		},
		{
			name:     "All the same",
			input:    []int{1, 1, 1, 1},
			expected: 1,
		},
		// Add more test cases if needed
	}

	for _, tt := range tests {
		t.Run(tt.name+"/removeDuplicates", func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input) // Copy the input slice to avoid modifying the test case data
			got := removeDuplicates(input)
			if got != tt.expected {
				t.Errorf("removeDuplicates() got = %v, want %v", got, tt.expected)
			}
		})

		t.Run(tt.name+"/removeDuplicates2", func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input) // Copy the input slice to avoid modifying the test case data
			got := removeDuplicates2(input)
			if got != tt.expected {
				t.Errorf("removeDuplicates2() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
