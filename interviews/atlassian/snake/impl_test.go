package snake

import (
	"testing"
)

func TestSnakeGame(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		food     [][]int
		moves    []string
		expected []int
	}{
		{
			name:     "TestGame1",
			width:    3,
			height:   2,
			food:     [][]int{{1, 2}, {0, 1}},
			moves:    []string{"R", "D", "R", "U", "L", "U"},
			expected: []int{0, 0, 1, 1, 2, -1},
		},
		{
			name:     "TestGame2",
			width:    3,
			height:   3,
			food:     [][]int{{2, 0}, {0, 0}, {2, 2}},
			moves:    []string{"D", "D", "R", "U", "U", "L", "D", "R", "R", "U", "L", "D"},
			expected: []int{0, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2},
		},
		// Additional test cases can be added here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := Constructor(tt.width, tt.height, tt.food)
			for i, move := range tt.moves {
				got := game.Move(move)
				if got != tt.expected[i] {
					t.Errorf("Move(%s) = %d; want %d", move, got, tt.expected[i])
				}
			}
		})
	}
}
