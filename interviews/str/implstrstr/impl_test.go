package implstrstr

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		input2 string
		want   int
	}{
		{"test 1", "sadbutsad", "sad", 0},
		{"test 2", "leetcode", "leeto", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := strStr(tt.input1, tt.input2)
			if tt.want != output {
				t.Errorf("strStr() = %v, want %v", output, tt.want)
			}
		})
	}
}
