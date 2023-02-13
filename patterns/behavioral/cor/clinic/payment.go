package clinic

type Payment struct {
	next Step
}

func (pa *Payment) execute(p *Patient) (resp string) {
	if p.paymentDone {
		p.record = pa.next.execute(p)
		return p.record + "Payment Done, "
	}

	p.paymentDone = true
	p.record = pa.next.execute(p)
	return p.record + "Reception getting money from patient, "
}

func (pa *Payment) setNext(next Step) {
	pa.next = next
}
