package maxfreqstr

import "testing"

func TestMaxFreqStr(t *testing.T) {
	tests := []struct {
		name       string
		inputStr   string
		maxLetters int
		minSize    int
		maxSize    int
		expected   int
	}{
		{
			name:       "Example 1",
			inputStr:   "aababcaab",
			maxLetters: 2,
			minSize:    3,
			maxSize:    4,
			expected:   2,
		},
		{
			name:       "Example 2",
			inputStr:   "aaaa",
			maxLetters: 1,
			minSize:    3,
			maxSize:    3,
			expected:   2,
		},
		{
			name:       "Example 3",
			inputStr:   "bbacbadadc",
			maxLetters: 2,
			minSize:    1,
			maxSize:    1,
			expected:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := maxFreqStr(tt.inputStr, tt.maxLetters, tt.minSize, tt.maxSize)

			if tt.expected != output {
				t.Errorf("maxFreStr() output is %v, expected output is %v", output, tt.expected)
			}
		})
	}
}
