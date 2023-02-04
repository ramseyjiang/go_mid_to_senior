package print

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdapter(t *testing.T) {
	msg := "Hello World!"

	adapter := OutputAdapter{OldSystem: &InfoLegacySystem{}, Msg: msg}
	assert.Equal(t, "Legacy System: Adapter: Hello World!", adapter.OutputStored())

	adapter = OutputAdapter{OldSystem: nil, Msg: msg}
	assert.Equal(t, "Hello World!", adapter.OutputStored())
}
