package bestbuysellstock

import "testing"

func TestBestBuySellStock(t *testing.T) {
	testCases := []struct {
		name   string
		inputs []int
		want   int
	}{
		{
			"test 1",
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			"test 2",
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			"test 3",
			[]int{1, 2},
			1,
		},
		{
			"test 4",
			[]int{2, 1, 2, 1, 0, 1, 2},
			2,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.inputs)
			if got != tt.want {
				t.Errorf("maxProfit(%v) got %v, want %v", tt.inputs, got, tt.want)
			}
		})
	}
}
