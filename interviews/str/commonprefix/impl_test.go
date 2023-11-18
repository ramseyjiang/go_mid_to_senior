package commonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{"test 1", []string{"flower", "flow", "flight"}, "fl"},
		{"test 2", []string{"dog", "racecar", "car"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := longestCommonPrefix(tt.input)
			if output != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", output, tt.want)
			}
		})
	}
}
