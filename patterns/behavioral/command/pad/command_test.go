package pad

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCommandExecute(t *testing.T) {
	pad := &Pad{}

	onCommand := &OnCommand{
		device: pad,
	}

	offCommand := &OffCommand{
		device: pad,
	}

	onButton := &Button{
		command: onCommand,
	}
	assert.Equal(t, "Turning pad on", onButton.press())

	offButton := &Button{
		command: offCommand,
	}
	assert.Equal(t, "Turning pad off", offButton.press())
}
