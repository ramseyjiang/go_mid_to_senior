package print

import "fmt"

type PastPrinter interface {
	Print(s string) string
}

type MyPastPrinter struct{}

func (l *MyPastPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Past Printer: %s", s)
	println(newMsg)
	return
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter PastPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}

	return
}
