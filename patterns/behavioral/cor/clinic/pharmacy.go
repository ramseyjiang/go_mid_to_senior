package clinic

type Pharmacy struct {
	next Step
}

func (ph *Pharmacy) execute(p *Patient) (resp string) {
	if p.pharmacyDone {
		return "Pharmacy already given to patient, "
	}

	p.pharmacyDone = true
	return "Pharmacy giving medicine to patient, "
}

func (ph *Pharmacy) setNext(next Step) {
	ph.next = next
}
