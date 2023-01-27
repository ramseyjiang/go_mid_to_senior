package pin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Check(t *testing.T) {
	var result bool
	testNumSlice := GeneratePins()

	for _, testNum := range testNumSlice {
		result = hasConsecutiveSame(testNum) && hasIncrementalConsecutive(testNum)
		// fmt.Println(testNum)
		assert.Equal(t, result, false)
	}
}
