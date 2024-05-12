package crawlerlogfolder

import (
	"testing"
)

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "Test 1",
			input:    []string{"d1/", "d2/", "../", "d21/", "./"},
			expected: 2,
		},
		{
			name:     "Test 2",
			input:    []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			expected: 3,
		},
		{
			name:     "Test 3",
			input:    []string{"d1/", "../", "../", "../"},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := MinOperations(tt.input)
			if tt.expected != output {
				t.Errorf("Output is %v, but expected is %v.", output, tt.expected)
			}
		})
	}
}
