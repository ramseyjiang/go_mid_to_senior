package office

type Printer interface {
	PrintFile() string
}

type Epson struct {
}

func (p *Epson) PrintFile() string {
	return "Printing by a EPSON Printer"
}

type Hp struct {
}

func (p *Hp) PrintFile() string {
	return "Printing by a HP Printer"
}
