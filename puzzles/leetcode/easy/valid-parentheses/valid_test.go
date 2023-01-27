package valid

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestValid(t *testing.T) {
	s1 := "()"
	wanted1 := isValid(s1)
	assert.Equal(t, wanted1, true)

	s2 := "()[]{}"
	wanted2 := isValid(s2)
	assert.Equal(t, wanted2, true)

	s3 := "(}"
	wanted3 := isValid(s3)
	assert.Equal(t, wanted3, false)
}
