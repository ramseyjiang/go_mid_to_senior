package clinic

type Appointment struct {
	next Step
}

func (a *Appointment) execute(p *Patient) (resp []string) {
	if a.next != nil {
		p.record = a.next.execute(p)
	}

	if p.bookDone {
		return append(p.record, "Patient appointment already done")
	}

	p.bookDone = true
	return append(p.record, "Reception appointment patient")
}

func (a *Appointment) setNext(next Step) {
	a.next = next
}
