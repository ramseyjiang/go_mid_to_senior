package office

// Computer is Abstraction interface.
type Computer interface {
	Print() string
	SetPrinter(Printer)
}

// Mac struct and all methods are concrete abstraction struct.
type Mac struct {
	printer Printer
}

func (m *Mac) Print() string {
	return "Print request for mac, " + m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// Windows struct and all methods are also concrete implementations.
type Windows struct {
	printer Printer
}

func (w *Windows) Print() string {
	return "Print request for windows, " + w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
