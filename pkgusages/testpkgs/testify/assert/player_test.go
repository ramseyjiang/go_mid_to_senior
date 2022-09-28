package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHadAGoodGame(t *testing.T) {
	tests := []struct {
		name     string
		stats    Stats
		goodGame bool
		wantErr  string
	}{
		{"sad path: invalid stats", Stats{Name: "Sam Cassell",
			Minutes:   34.1,
			Points:    -19,
			Assists:   8,
			Turnovers: -4,
			Rebounds:  11,
		}, false, "stat lines cannot be negative",
		},
		{"happy path: good game", Stats{Name: "Dejounte Murray",
			Minutes:   34.1,
			Points:    19,
			Assists:   8,
			Turnovers: 4,
			Rebounds:  11,
		}, true, ""},
	}
	for _, tt := range tests {
		isAGoodGame, err := hadAGoodGame(tt.stats)
		if tt.wantErr != "" {
			assert.Contains(t, err.Error(), tt.wantErr)
		} else {
			assert.Equal(t, tt.goodGame, isAGoodGame)
		}
	}
}
