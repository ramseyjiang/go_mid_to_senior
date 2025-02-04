package findmaxchainlen

import "testing"

func TestFindMaxChainLength(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Basic case",
			input:    "ecbdca",
			expected: 3, // "bcd" and "cbd" are longest valid substrings
		},
		{
			name:     "All increasing order",
			input:    "abcdef",
			expected: 6, // Whole string is valid
		},
		{
			name:     "All decreasing order",
			input:    "fedcba",
			expected: 0, // No valid substrings as start >= end
		},
		{
			name:     "Single character",
			input:    "a",
			expected: 0, // Less than 2 characters, should return 0
		},
		{
			name:     "Two characters valid",
			input:    "ac",
			expected: 2, // "ac" is valid as 'a' < 'c'
		},
		{
			name:     "Two characters invalid",
			input:    "ca",
			expected: 0, // 'c' is not < 'a'
		},
		{
			name:     "Duplicate characters with gaps",
			input:    "abcabc",
			expected: 6, // "abcabc" covers the longest valid substring
		},
		{
			name:     "Only one valid pair",
			input:    "za",
			expected: 0, // No valid substring
		},
		{
			name:     "Long string with random characters",
			input:    "xhijklmnaopqrstuvw",
			expected: 17, // "hijklmnopqrstuvw" is the longest valid substring
		},
		{
			name:     "Large input",
			input:    "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			expected: 52, // Whole alphabet is a valid substring
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findMaxChainLength(tt.input)
			if actual != tt.expected {
				t.Errorf("%s: expected %d, got %d", tt.name, tt.expected, actual)
			}
		})
	}
}
