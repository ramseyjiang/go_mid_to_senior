package reorganisestr

import (
	"testing"
)

func TestReorganizeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string // Any valid permutation or ""
	}{
		{"Example 1", "aab", "aba"},
		{"Example 2 (Impossible Case)", "aaab", ""}, // Must return ""
		{"Single Character", "a", "a"},
		{"Two Different Characters", "ab", "ab"},
		{"Multiple Characters", "aaabbc", "ababac"},
		{"Balanced Characters", "aabbcc", "abcabc"},
		{"Long Impossible Case", "aaaaaabb", ""}, // Too many 'a's to rearrange
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reorganizeString(tt.input)
			if !isValidReorganization(tt.input, result) {
				t.Errorf("Test %s failed: Expected valid reorganization but got %s", tt.name, result)
			}
		})
	}
}

// Helper function to check if a string is a valid reorganization of the input
func isValidReorganization(s, result string) bool {
	if result == "" {
		// If the function returns "", it must be because it's impossible to reorganize
		// We check if the most frequent character appears more than (n+1)/2 times
		freq := make(map[rune]int)
		n := len(s)
		maxFreq := 0
		for _, c := range s {
			freq[c]++
			if freq[c] > maxFreq {
				maxFreq = freq[c]
			}
		}
		return maxFreq > (n+1)/2 // Must return "" if a valid arrangement is impossible
	}

	// Check adjacent characters are not the same
	for i := 1; i < len(result); i++ {
		if result[i] == result[i-1] {
			return false
		}
	}

	// Check if the result is a valid permutation of the input
	freqS := make(map[rune]int)
	freqR := make(map[rune]int)
	for _, c := range s {
		freqS[c]++
	}
	for _, c := range result {
		freqR[c]++
	}
	for k, v := range freqS {
		if freqR[k] != v {
			return false
		}
	}

	return true
}
