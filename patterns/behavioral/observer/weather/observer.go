package weather

// Observer interface
type Observer interface {
	Update(subject Subject)
	GetTemp() float64
}

// PhoneDisplay is a Concrete Observer
type PhoneDisplay struct {
	temperature float64
}

func (p *PhoneDisplay) Update(subject Subject) {
	if w, ok := subject.(*Station); ok {
		p.temperature = w.temperature
	}
}

func (p *PhoneDisplay) GetTemp() float64 {
	return p.temperature
}

// PCDisplay is a Concrete Observer
type PCDisplay struct {
	temperature float64
}

func (p *PCDisplay) Update(subject Subject) {
	if w, ok := subject.(*Station); ok {
		p.temperature = w.temperature
	}
}

func (p *PCDisplay) GetTemp() float64 {
	return p.temperature
}
