package clinic

type Pharmacy struct {
	next Handler
}

func (ph *Pharmacy) execute(p *Patient) (resp []string) {
	if ph.next != nil {
		p.record = ph.next.execute(p)
	}

	if p.pharmacyDone {
		return append(p.record, "Pharmacy already given to patient")
	}

	p.pharmacyDone = true
	return append(p.record, "Pharmacy giving medicine to patient")
}

func (ph *Pharmacy) setNext(next Handler) {
	ph.next = next
}
