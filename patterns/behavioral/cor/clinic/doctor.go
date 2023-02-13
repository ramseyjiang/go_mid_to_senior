package clinic

type Doctor struct {
	next Step
}

func (d *Doctor) execute(p *Patient) (resp string) {
	if p.doctorCheckUpDone {
		p.record = d.next.execute(p)
		return p.record + "Doctor checkup already done, "
	}

	p.doctorCheckUpDone = true
	p.record = d.next.execute(p)
	return p.record + "Doctor checking patient, "
}

func (d *Doctor) setNext(next Step) {
	d.next = next
}
