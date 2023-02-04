package office

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPrint(t *testing.T) {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

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
