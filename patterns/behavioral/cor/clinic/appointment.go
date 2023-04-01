package clinic

type Appointment struct {
	next Handler
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

func (a *Appointment) setNext(next Handler) {
	a.next = next
}

func createHandlerChain() Handler {
	pharmacy := &Pharmacy{}
	payment := &Payment{}
	doctor := &Doctor{}
	appointment := &Appointment{}

	// Set next for doctor step
	payment.setNext(pharmacy)

	// Set next for payment step
	doctor.setNext(payment)

	// Set next for pharmacy step
	appointment.setNext(doctor)

	return appointment
}
