package strtoint

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"test 1", "42", 42},
		{"test 2", "   -42", -42},
		{"test 3", "4193 with words", 4193},
		{"test 4", "words and 987", 0},
		{"test 5", "-91283472332", -2147483648},
		{"test 6", "  -0012a42", -12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := myAtoi(tt.input)
			if output != tt.want {
				t.Errorf("myAtoi() = %v, want %v", output, tt.want)
			}
		})
	}
}
