package assert

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestHadAGoodGame(t *testing.T) {
	tests := []struct {
		name          string
		stats         Stats
		goodGame      bool
		errorExpected bool
		errorMessage  string
	}{
		{"sad path: invalid stats",
			Stats{Name: "Sam Cassell",
				Minutes:   34.1,
				Points:    -19,
				Assists:   8,
				Turnovers: -4,
				Rebounds:  11,
			},
			false,
			true,
			"stat lines cannot be negative",
		},
		{"happy path: good game",
			Stats{Name: "Dejounte Murray",
				Minutes:   34.1,
				Points:    19,
				Assists:   8,
				Turnovers: 4,
				Rebounds:  11,
			},
			true,
			false,
			""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			isAGoodGame, err := hadAGoodGame(tt.stats)

			if tt.errorExpected {
				// Assert
				require.Error(t, err)
				assert.Equal(t, tt.errorMessage, err.Error())
			}
			assert.Equal(t, tt.goodGame, isAGoodGame)
		})
	}
}
