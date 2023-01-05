package redispkg

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_getRedisValues(t *testing.T) {
	wanted := [3]string{"customise-value", "customise-value", ""}
	val, val1, val2 := getRedisValues()
	assert.Equal(t, val, wanted[0])
	assert.Equal(t, val1, wanted[1])
	assert.Equal(t, val2, wanted[2])
}
