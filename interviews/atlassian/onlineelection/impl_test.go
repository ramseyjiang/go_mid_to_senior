package onlineelection

import (
	"fmt"
	"testing"
)

func TestTopVotedCandidate(t *testing.T) {
	persons := []int{0, 1, 1, 0, 0, 1, 0}
	times := []int{0, 5, 10, 15, 20, 25, 30}
	tvc := NewOnlineElection(persons, times)

	tests := []struct {
		t        int
		expected int
	}{
		{3, 0},
		{12, 1},
		{25, 1},
		{15, 0},
		{24, 0},
		{8, 1},
		// Additional test cases here
	}

	for _, tt := range tests {
		t.Run("Query at time "+fmt.Sprint(tt.t), func(t *testing.T) {
			got := tvc.Query(tt.t)
			if got != tt.expected {
				t.Errorf("Q(%d) = %d, want %d", tt.t, got, tt.expected)
			}
		})
	}
}
