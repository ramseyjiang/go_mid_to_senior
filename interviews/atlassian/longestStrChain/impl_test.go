package longestStrChain

import "testing"

func TestLongestStrChain(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected int
	}{
		{
			name:     "Example 1",
			words:    []string{"a", "b", "ba", "bca", "bda", "bdca"},
			expected: 4,
		},
		{
			name:     "Example 2",
			words:    []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
			expected: 5,
		},
		{
			name:     "Example 3",
			words:    []string{"abcd", "dbqca"},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := longestStrChan(tt.words)
			if output != tt.expected {
				t.Errorf("longestStrChan() output is %v, expected is %v", output, tt.expected)
			}
		})
	}
}
