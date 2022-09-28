package assert

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalAndPrint(t *testing.T) {
	t.Run("testing unmarshalAndPrint()", func(t *testing.T) {
		err := unmarshalAndPrint(strings.NewReader(`[{"name": "Dubi Gal", "age": 900}]`))
		assert.Nil(t, err)
	})
}
