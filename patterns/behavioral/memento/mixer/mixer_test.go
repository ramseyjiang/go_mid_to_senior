package mixer

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// get a variable with a Facade pattern.
// Zero-value initialization will give us zero-valued originator and caretaker objects.
func TestMixer(t *testing.T) {
	m := MementoFacade{}

	// create a Volume value with Volume(4), it uses parentheses.
	// The Volume type does not have any inner field like structs, so we cannot use curly braces to set its value.
	m.SaveSettings(Volume(4))

	// create a Mute value with Mute(false), it also uses parentheses. The same reason with above.
	m.SaveSettings(Mute(false))

	// anonymous function
	selection := func(c Command) any {
		switch cast := c.(type) {
		case Volume:
			return cast
		case Mute:
			return cast
		default:
			return ""
		}
	}

	var expected1 Mute = false
	assert.Equal(t, expected1, selection(m.RestoreSettings()))

	var expected2 Volume = 4
	assert.Equal(t, expected2, selection(m.RestoreSettings()))
}
