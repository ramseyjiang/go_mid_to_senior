package isanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		input2 string
		want   bool
	}{
		{"test 1", "anagram", "nagaram", true},
		{"test 2", "rat", "car", false},
		{"test 3", "ab", "a", false},
		{"test 4", "aacc", "ccac", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := isAnagram(tt.input1, tt.input2)
			if output != tt.want {
				t.Errorf("isAnagram() = %v, want %v", output, tt.want)
			}
		})
	}
}
