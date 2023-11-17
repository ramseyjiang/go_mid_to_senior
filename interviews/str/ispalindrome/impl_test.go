package ispalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"test 1", "A man, a plan, a canal: Panama", true},
		{"test 2", "race a car", false},
		{"test 3", " ", true},
		{"test 4", "0P", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := isPalindrome(tt.input)
			if output != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", output, tt.want)
			}
		})
	}
}
