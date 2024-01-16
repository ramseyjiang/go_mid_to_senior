package cherrypickup

import "testing"

func TestCherryPickup(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "Example 1",
			grid:     [][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}},
			expected: 5,
		},
		{
			name:     "Example 2",
			grid:     [][]int{{1, 1, -1}, {1, -1, 1}, {-1, 1, 1}},
			expected: 0,
		},
		// Additional test cases can be added here.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cherryPickup(tt.grid)
			if got != tt.expected {
				t.Errorf("cherryPickup(%v) = %d, want %d", tt.grid, got, tt.expected)
			}
		})
	}
}
