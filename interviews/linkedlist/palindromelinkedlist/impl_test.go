package palindromelinkedlist

import "testing"

var tests = []struct {
	name     string
	input    []int
	expected bool
}{
	{
		name:     "Test 1",
		input:    []int{1, 2, 2, 1},
		expected: true,
	},
	{
		name:     "Test 2",
		input:    []int{1, 2},
		expected: false,
	},
	{
		name:     "Test 3",
		input:    []int{},
		expected: true,
	},
	{
		name:     "Test 4",
		input:    []int{1, 0, 1},
		expected: true,
	},
	{
		name:     "Test 5",
		input:    []int{1, 0, 2, 1},
		expected: false,
	},
}

func TestIsPalindromeIterative(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			result := isPalindromeIterative(head)
			if result != tt.expected {
				t.Errorf("isPalindromeIterative() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsPalindromeRecursive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			result := isPalindromeRecursive(head)
			if result != tt.expected {
				t.Errorf("isPalindromeRecursive() = %v, want %v", result, tt.expected)
			}
		})
	}
}
