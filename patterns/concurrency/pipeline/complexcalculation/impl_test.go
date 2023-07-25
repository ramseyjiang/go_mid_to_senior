package complexcalculation

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}

	var res int
	for _, test := range tableTest {
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Fatal()
		}
		assert.Equal(t, res, test[1])
	}
}
