package office

type Computer interface {
	Print() string
	SetPrinter(Printer)
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() string {
	return "Print request for mac, " + m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() string {
	return "Print request for windows, " + w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
