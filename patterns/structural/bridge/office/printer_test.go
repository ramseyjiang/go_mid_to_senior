package office

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPrint(t *testing.T) {
	// all client code are below, except to assert parts.
	hpPrinter := &Hp{}
	assert.Equal(t, "Printing by an HP Printer", hpPrinter.PrintFile())

	epsonPrinter := &Epson{}
	assert.Equal(t, "Printing by an EPSON Printer", epsonPrinter.PrintFile())

	mac := &Mac{}
	mac.SetPrinter(hpPrinter)
	assert.Equal(t, "Print request for mac, Printing by an HP Printer", mac.Print())

	mac.SetPrinter(epsonPrinter)
	assert.Equal(t, "Print request for mac, Printing by an EPSON Printer", mac.Print())

	windows := &Windows{}
	windows.SetPrinter(hpPrinter)
	assert.Equal(t, "Print request for windows, Printing by an HP Printer", windows.Print())

	windows.SetPrinter(epsonPrinter)
	assert.Equal(t, "Print request for windows, Printing by an EPSON Printer", windows.Print())
}
