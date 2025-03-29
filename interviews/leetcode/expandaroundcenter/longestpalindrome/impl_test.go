package longestPalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1",
			input:    "babad",
			expected: "bab", // "aba" is also acceptable
		},
		{
			name:     "Example 2",
			input:    "cbbd",
			expected: "bb",
		},
		{
			name:     "Single Character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Entire String is Palindromic",
			input:    "racecar",
			expected: "racecar",
		},
		{
			name:     "Mixed Characters",
			input:    "abcba",
			expected: "abcba",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("longestPalindrome(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
