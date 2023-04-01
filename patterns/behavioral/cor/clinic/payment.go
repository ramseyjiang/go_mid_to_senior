package clinic

type Payment struct {
	next Handler
}

func (pa *Payment) execute(p *Patient) (resp []string) {
	if pa.next != nil {
		p.record = pa.next.execute(p)
	}

	if p.paymentDone {
		return append(p.record, "Payment Done")
	}

	p.paymentDone = true
	return append(p.record, "Reception getting money from patient")
}

func (pa *Payment) setNext(next Handler) {
	pa.next = next
}
