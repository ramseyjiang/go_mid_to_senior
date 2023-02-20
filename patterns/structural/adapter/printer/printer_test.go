package print

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdapter(t *testing.T) {
	msg := "Hello World!"

	// The Client orchestrates the adapter by calling the adaptee’s method indirectly.
	adapter := PrinterAdapter{Legacy: &LegacyPrinterImpl{}, Msg: msg}
	assert.Equal(t, "Legacy Printer: Adapter: Hello World!", adapter.PrintMessage())

	adapter = PrinterAdapter{Legacy: nil, Msg: msg}
	assert.Equal(t, "Hello World!", adapter.PrintMessage())
}
