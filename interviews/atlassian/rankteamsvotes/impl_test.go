package rankteamsvotes

import "testing"

func TestRankVotes(t *testing.T) {
	tests := []struct {
		Name        string
		InputSlice  []string
		ExpectedStr string
	}{
		{
			Name:        "Example 1",
			InputSlice:  []string{"ABC", "ACB", "ABC", "ACB", "ACB"},
			ExpectedStr: "ACB",
		},
		{
			Name:        "Example 2",
			InputSlice:  []string{"WXYZ", "XYZW"},
			ExpectedStr: "XWYZ",
		},
		{
			Name:        "Example 3",
			InputSlice:  []string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"},
			ExpectedStr: "ZMNAGUEDSJYLBOPHRQICWFXTVK",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			output := RankVotes(tt.InputSlice)
			if tt.ExpectedStr != output {
				t.Errorf("RankVotes result is %v, but the expected output is %v", output, tt.ExpectedStr)
			}
		})
	}
}
