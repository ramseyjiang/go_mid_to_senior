package print

import "fmt"

// ModernPrinter is the interface that Client uses, the ModernPrinter is the target.
type ModernPrinter interface {
	PrintMessage() string
}

// LegacyPrinter is an existing type with a specific interface that needs to be adapted.
type LegacyPrinter interface {
	Print(s string) string
}

// LegacyPrinterImpl is an implementation of the LegacyPrinter interface, the LegacyPrinterImpl is the Adaptee.
type LegacyPrinterImpl struct{}

// Print is a method that implements the LegacyPrinter interface and modifies the passed string by prefixing the text "Legacy Printer:"
func (l *LegacyPrinterImpl) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s", s)
	return
}

// PrinterAdapter implements the ModernPrinter interface by using an instance of the LegacyPrinter struct, so it allows the LegacyPrinter to be used
type PrinterAdapter struct {
	Legacy LegacyPrinter
	Msg    string
}

// PrintMessage method of the ModernPrinter interface; this method doesn't accept any argument and must return the modified string.
// The PrinterAdapter implements the above target interface. It is an adapter between LegacyPrinter and ModernPrinter.
func (p *PrinterAdapter) PrintMessage() (newMsg string) {
	if p.Legacy != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.Legacy.Print(newMsg)
	} else {
		newMsg = p.Msg
	}

	return
}
