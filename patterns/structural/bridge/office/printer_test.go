package office

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPrint(t *testing.T) {
	// all client code are below, except to assert parts.
	hpPrinter := &Hp{}
	assert.Equal(t, "Printing by a HP Printer", hpPrinter.PrintFile())

	epsonPrinter := &Epson{}
	assert.Equal(t, "Printing by a EPSON Printer", epsonPrinter.PrintFile())

	macComputer := &Mac{}
	macComputer.SetPrinter(hpPrinter)
	assert.Equal(t, "Print request for mac, Printing by a HP Printer", macComputer.Print())

	macComputer.SetPrinter(epsonPrinter)
	assert.Equal(t, "Print request for mac, Printing by a EPSON Printer", macComputer.Print())

	winComputer := &Windows{}
	winComputer.SetPrinter(hpPrinter)
	assert.Equal(t, "Print request for windows, Printing by a HP Printer", winComputer.Print())

	winComputer.SetPrinter(epsonPrinter)
	assert.Equal(t, "Print request for windows, Printing by a EPSON Printer", winComputer.Print())
}
