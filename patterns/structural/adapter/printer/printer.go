package print

import "fmt"

type LegacySystem interface {
	Output(s string) string
}

type InfoLegacySystem struct{}

// Output is a method that implements the LegacySystem interface and modifies the passed string by prefixing the text "Legacy System:"
func (l *InfoLegacySystem) Output(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy System: %s", s)
	println(newMsg)
	return
}

type NewSystem interface {
	OutputStored() string
}

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
