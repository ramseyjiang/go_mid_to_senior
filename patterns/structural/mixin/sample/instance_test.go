package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMixin(t *testing.T) {
	mixinReceiver := NewReceiver()
	assert.Equal(t, "This is the Mixin Method", mixinReceiver.MixinMethod())
	assert.Equal(t, "This is the Receiver Method", mixinReceiver.ReceiverMethod())
}
