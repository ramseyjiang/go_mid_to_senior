package office

// Printer is Implementation interface
type Printer interface {
	PrintFile() string
}

// Epson struct and the method with Epson are Concrete implementations.
type Epson struct {
}

func (p *Epson) PrintFile() string {
	return "Printing by an EPSON Printer"
}

// Hp struct and the method with Hp are Concrete implementations.
type Hp struct {
}

func (p *Hp) PrintFile() string {
	return "Printing by an HP Printer"
}
