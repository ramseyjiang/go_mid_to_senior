package firstunique

import (
	"testing"
)

func TestFirstUniqueStr(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"test 1", "leetcode", 0},
		{"test 2", "loveleetcode", 2},
		{"test 3", "aabb", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := firstUniqueChar(tt.input)
			if output != tt.want {
				t.Errorf("firstUniqueChar() = %v, want %v", output, tt.want)
			}
		})
	}
}
