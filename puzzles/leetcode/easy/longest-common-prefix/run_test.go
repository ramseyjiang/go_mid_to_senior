package longest_common_prefix

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLongestCommonPrefix(t *testing.T) {
	str1 := []string{"flower", "flow", "flight"}
	wanted1 := longestCommonPrefix(str1)
	expected1 := "fl"
	assert.Equal(t, wanted1, expected1)

	str2 := []string{"dog", "race car", "car"}
	wanted2 := longestCommonPrefix(str2)
	expected2 := ""
	assert.Equal(t, wanted2, expected2)

	str3 := []string{"reflower", "flow", "flight"}
	wanted3 := longestCommonPrefix(str3)
	expected3 := ""
	assert.Equal(t, wanted3, expected3)
}
