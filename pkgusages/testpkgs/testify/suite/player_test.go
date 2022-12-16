package suite

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/suite"
)

type GameTestSuite struct {
	suite.Suite
}

func (suite *GameTestSuite) SetupSuite() {
	// create relevant resources
	fmt.Println("SETUP")
}

func (suite *GameTestSuite) BeforeTest(_, _ string) {
	// execute code before test starts
	fmt.Println("BEFORE")
}

func (suite *GameTestSuite) AfterTest(_, _ string) {
	// execute code after test finishes
	fmt.Println("AFTER")
}

func TestGameTestSuite(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}

// notice: suite.Require().Equal() usage is different from assert.Equal().
func (suite *GameTestSuite) TestHadAGoodGame() {
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
		suite.T().Run("test hadAGoodGame(): "+tt.name, func(t *testing.T) {
			// Act
			isAGoodGame, err := hadAGoodGame(tt.stats)

			if tt.errorExpected {
				// Assert
				require.Error(t, err)
				suite.Require().Equal(tt.errorMessage, err.Error())
			}
			suite.Require().Equal(tt.goodGame, isAGoodGame)
		})
	}
}
