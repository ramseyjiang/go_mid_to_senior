package office

// Computer is Abstraction interface.
type Computer interface {
	Print() string
	SetPrinter(Printer)
}

// Mac and all method within mac are Refined abstraction
type Mac struct {
	printer Printer
}

func (m *Mac) Print() string {
	return "Print request for mac, " + m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// Windows and all method within Windows are Refined abstraction
type Windows struct {
	printer Printer
}

func (w *Windows) Print() string {
	return "Print request for windows, " + w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
