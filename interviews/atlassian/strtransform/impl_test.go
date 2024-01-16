package strtransform

import "testing"

func TestTransformString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		k        int64
		expected int
	}{
		{
			name:     "Example 1",
			s:        "abcd",
			t:        "cdab",
			k:        2,
			expected: 2,
		},
		{
			name:     "Example 2",
			s:        "ababab",
			t:        "ababab",
			k:        1,
			expected: 2,
		},
		// Additional test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfWays(tt.s, tt.t, tt.k)
			if got != tt.expected {
				t.Errorf("numberOfWays(%s, %s, %d) = %d, want %d", tt.s, tt.t, tt.k, got, tt.expected)
			}
		})
	}
}
