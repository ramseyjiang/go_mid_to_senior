package clinic

type Doctor struct {
	next Step
}

func (d *Doctor) execute(p *Patient) (resp string) {
	if d.next != nil {
		p.record = d.next.execute(p)
	}

	if p.doctorCheckUpDone {
		return p.record + "Doctor checkup already done, "
	}

	p.doctorCheckUpDone = true
	return p.record + "Doctor checking patient, "
}

func (d *Doctor) setNext(next Step) {
	d.next = next
}
