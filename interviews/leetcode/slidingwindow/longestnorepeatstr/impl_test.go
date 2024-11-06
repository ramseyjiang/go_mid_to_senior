package longestnorepeatstr

import "testing"

func TestFindLongestNoRepeatStr(t *testing.T) {
	tests := []struct {
		name     string
		strs     string
		expected int
	}{
		{
			"Test 1",
			"abcabcbb",
			3,
		},
		{
			"Test 2",
			"bbbbb",
			1,
		},
		{
			"Test 3",
			"pwwkew",
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findLongestNoRepeatStr(tt.strs)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
