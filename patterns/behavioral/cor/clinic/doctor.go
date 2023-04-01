package clinic

type Doctor struct {
	next Handler
}

func (d *Doctor) execute(p *Patient) (resp []string) {
	if d.next != nil {
		p.record = d.next.execute(p)
	}

	if p.doctorCheckUpDone {
		return append(p.record, "Doctor checkup already done")
	}

	p.doctorCheckUpDone = true
	return append(p.record, "Doctor checking patient")
}

func (d *Doctor) setNext(next Handler) {
	d.next = next
}
