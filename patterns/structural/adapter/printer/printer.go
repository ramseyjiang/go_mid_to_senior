package print

import "fmt"

// NewSystem is the interface that Client uses
type NewSystem interface {
	OutputStored() string
}

// LegacySystem is a third-party system with a different interface
type LegacySystem interface {
	Output(s string) string
}

// InfoLegacySystem is an existing struct with the LegacySystem interface
type InfoLegacySystem struct{}

// Output is a method that implements the LegacySystem interface and modifies the passed string by prefixing the text "Legacy System:"
func (l *InfoLegacySystem) Output(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy System: %s", s)
	println(newMsg)
	return
}

// OutputAdapter implements the NewSystem interface by using an instance of the LegacySystem struct, so it allows the LegacySystem to be used
type OutputAdapter struct {
	OldSystem LegacySystem
	Msg       string
}

// OutputStored method of the NewSystem interface; this method doesn't accept any argument and must return the modified string.
// It is an adapter between LegacySystem and NewSystem.
func (p *OutputAdapter) OutputStored() (newMsg string) {
	if p.OldSystem != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldSystem.Output(newMsg)
	} else {
		newMsg = p.Msg
	}

	return
}
