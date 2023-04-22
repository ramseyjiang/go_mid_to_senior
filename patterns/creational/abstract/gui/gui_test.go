package gui

import (
	"runtime"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMacGUI(t *testing.T) {
	var factory FactoryGUI

	// Choose the concrete factory based on the platform or configuration.
	if runtime.GOOS == "darwin" {
		factory = &MacOSFactory{}
	} else {
		factory = &WindowsFactory{}
	}

	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	assert.Equal(t, button.Click(), "MacOS button clicked")
	assert.Equal(t, checkbox.Check(), "MacOS checkbox checked")
}
