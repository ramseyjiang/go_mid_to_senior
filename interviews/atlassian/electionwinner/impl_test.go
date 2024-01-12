package electionwinner

import "testing"

func TestWinner(t *testing.T) {
	tests := []struct {
		name        string
		inputArr    []string
		expectedStr string
		expectedNum int
	}{
		{
			name:        "Example 1",
			inputArr:    []string{"john", "johnny", "jackie", "johnny", "john", "jackie", "jamie", "jamie", "john", "johnny", "jamie", "johnny", "john"},
			expectedStr: "john",
			expectedNum: 4,
		},
		{
			name:        "Example 2",
			inputArr:    []string{"andy", "blake", "clark"},
			expectedStr: "andy",
			expectedNum: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputStr, outputNum := winner(tt.inputArr)
			if outputNum != tt.expectedNum || outputStr != tt.expectedStr {
				t.Errorf("winner() output string is %v and the number is %v, not as expected", outputNum, outputNum)
			}
		})
	}
}
