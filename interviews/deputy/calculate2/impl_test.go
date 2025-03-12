package calculate2

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "basic operations",
			s:    "3+2*2",
			want: 7,
		},
		{
			name: "division truncation",
			s:    " 3/2 ",
			want: 1,
		},
		{
			name: "mixed operations with spaces",
			s:    " 3+5 / 2 * 4-6 ",
			want: 3 + 2*4 - 6, // = 3 + 8 -6 = 5
		},
		{
			name: "negative result",
			s:    "1-1-1",
			want: -1,
		},
		{
			name: "complex precedence",
			s:    "1+2*5/3+6/4*2",
			want: 1 + (10 / 3) + (6 / 4 * 2), // = 1+3+3 =7
		},
		{
			name: "single number",
			s:    "  42  ",
			want: 42,
		},
		{
			name: "leading negative",
			s:    "-42 + 10", // 注意：原代码需要支持负数开头
			want: -32,
		},
		{
			name: "division order",
			s:    "14-3/2",
			want: 14 - 1, // =13
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.s); got != tt.want {
				t.Errorf("calculate(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
