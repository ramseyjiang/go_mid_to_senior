package print

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdapter(t *testing.T) {
	msg := "Hello World!"

	adapter := PrinterAdapter{OldPrinter: &MyPastPrinter{}, Msg: msg}
	assert.Equal(t, "Past Printer: Adapter: Hello World!", adapter.PrintStored())

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	assert.Equal(t, "Hello World!", adapter.PrintStored())
}
