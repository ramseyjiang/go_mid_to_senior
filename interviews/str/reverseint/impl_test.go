package reverseint

import "testing"

func TestReverseInt(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"test 1", 123, 321},
		{"test 2", -123, -321},
		{"test 3", 120, 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := reverseInt(tt.input)
			if output != tt.want {
				t.Errorf("reverseInt() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}
