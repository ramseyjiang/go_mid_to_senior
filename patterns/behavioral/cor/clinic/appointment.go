package clinic

type Appointment struct {
	next Step
}

func (a *Appointment) execute(p *Patient) (resp string) {
	if p.bookDone {
		p.record = a.next.execute(p)
		return p.record + "Patient appointment already done"
	}

	p.bookDone = true
	p.record = a.next.execute(p)
	return p.record + "Reception appointment patient"
}

func (a *Appointment) setNext(next Step) {
	a.next = next
}
