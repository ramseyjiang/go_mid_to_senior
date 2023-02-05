package office

// Printer is Implementation interface
type Printer interface {
	PrintFile() string
}

// Epson and method with Epson are Concrete implementation
type Epson struct {
}

func (p *Epson) PrintFile() string {
	return "Printing by a EPSON Printer"
}

// Hp and method with Hp are Concrete implementation
type Hp struct {
}

func (p *Hp) PrintFile() string {
	return "Printing by a HP Printer"
}
