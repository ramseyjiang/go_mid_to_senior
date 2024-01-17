package majorityelement

import (
	"testing"
)

func TestFindMajority(t *testing.T) {
	tests := []struct {
		name     string
		inputNum int
		inputArr []int
		expected int
	}{
		{name: "Example 1", inputNum: 3, inputArr: []int{1, 2, 3}, expected: -1},
		{name: "Example 2", inputNum: 5, inputArr: []int{3, 1, 3, 3, 2}, expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FindMajority(tt.inputNum, tt.inputArr)
			if output != tt.expected {
				t.Errorf("FindMajority() output is %v, expected is %v", output, tt.expected)
			}
		})
	}
}
