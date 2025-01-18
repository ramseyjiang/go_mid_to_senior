package happynumber

import "testing"

func TestIsHappy(t *testing.T) {
	testCases := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "Happy number",
			input: 19,
			want:  true,
		},
		{
			name:  "Unhappy number",
			input: 2,
			want:  false,
		},
		{
			name:  "Another happy number",
			input: 7,
			want:  true,
		},
		{
			name:  "Edge case: 1",
			input: 1,
			want:  true,
		},
		{
			name:  "Another unhappy number",
			input: 4,
			want:  false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := isHappy(tt.input)
			if got != tt.want {
				t.Errorf("isHappy(%d) got %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
