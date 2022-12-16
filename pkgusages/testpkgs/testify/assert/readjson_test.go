package assert

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalAndPrint(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Act
		err := unmarshalAndPrint(strings.NewReader(`[{"name": "Dubi Gal", "age": 90}]`))
		// Assert
		assert.NoError(t, err)
	})

	t.Run("sad path", func(t *testing.T) {
		// Act
		err := unmarshalAndPrint(strings.NewReader(`{"name": "Dubi Gal", "age": 90}`))
		// Assert
		assert.Error(t, err)
		assert.Equal(t, "json format error", err.Error())
	})
}
