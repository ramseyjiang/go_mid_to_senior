package twotestways

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPackDoughnutsBoxTableTests(t *testing.T) {
	// Anonymous struct of test cases
	tests := []struct {
		name                           string
		boxCapacity                    int
		errorExpected                  bool
		errorMessage                   string
		items                          []string
		expectedNumOfDoughnutsInTheBox int
	}{
		{
			name:                           "Filling the box with tasty doughnuts",
			boxCapacity:                    4,
			errorExpected:                  false,
			errorMessage:                   "",
			items:                          []string{"Sri Lankan Cinnamon Sugar", "Mocha Tea", "Home Made Raspberry Jam", "Lime & Coconut (ve)"},
			expectedNumOfDoughnutsInTheBox: 4,
		},
		{
			name:                           "Attempt to fill the box with too many doughnuts",
			boxCapacity:                    4,
			errorExpected:                  true,
			errorMessage:                   "failed to put 5 doughnuts in the box, it's only has 4 doughnuts capacity",
			items:                          []string{"Sri Lankan Cinnamon Sugar", "Mocha Tea", "Home Made Raspberry Jam", "Lime & Coconut (ve)", "Lime & Coconut (ve)"},
			expectedNumOfDoughnutsInTheBox: 0,
		},
		{
			name:                           "Attempt to put a giant chocolate cookie into the box",
			boxCapacity:                    2,
			errorExpected:                  true,
			errorMessage:                   "the following items cannot be placed into the box: [Giant Chocolate Cookie]",
			items:                          []string{"Sri Lankan Cinnamon Sugar", "Giant Chocolate Cookie"},
			expectedNumOfDoughnutsInTheBox: 1,
		},
	}

	for _, tc := range tests {
		// each test case from  table above run as a subtest
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			box := newDoughnutsBox(tc.boxCapacity)

			// Act
			numOfDoughnutsInTheBox, err := box.pack(tc.items)
			// fmt.Println(err, tc.name) // In the first test, err equal to nil, if output err.Error(), it will have an error during it is running.

			// Assert
			if tc.errorExpected {
				require.Error(t, err)
				// fmt.Println(err.Error(), tc.name) // When tests have err, err.Error() output is the same with err in this case.
				assert.Equal(t, tc.errorMessage, err.Error())
			}
			assert.Equal(t, tc.expectedNumOfDoughnutsInTheBox, numOfDoughnutsInTheBox)
		})
	}
}
