package divisorgame

import "testing"

func TestDivisorGame(t *testing.T) {
	testCases := []struct {
		name   string
		n      int
		expect bool
	}{
		{"n is even", 2, true},
		{"n is odd", 3, false},
		{"n is even and large", 1000, true},
		{"n is odd and large", 999, false},
		{"n is 1", 1, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := divisorGame(tc.n)
			if result != tc.expect {
				t.Errorf("divisorGame(%d) = %v, expected %v", tc.n, result, tc.expect)
			}
		})
	}
}
