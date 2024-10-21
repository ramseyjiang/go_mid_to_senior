package longestsubstrnorepeat

import "testing"

func TestFindLongestSubstr(t *testing.T) {
	tests := []struct {
		name     string
		inputStr string
		want     int
	}{
		{
			"test 1",
			"abcabcbb",
			3,
		},
		{
			"test 2",
			"bbbbb",
			1,
		},
		{
			"test 3",
			"pwwkew",
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLongestSubstr(tt.inputStr)
			if got != tt.want {
				t.Errorf("findLongestSubstr(%v) = %v, want %v", tt.inputStr, got, tt.want)
			}
		})
	}
}
