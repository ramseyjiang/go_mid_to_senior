package pad

// Device is the Receiver interface
type Device interface {
	on() string
	off() string
}

// Pad is a Concrete receiver
type Pad struct {
	isRunning bool
}

func (p *Pad) on() string {
	p.isRunning = true
	return "Turning pad on"
}

func (p *Pad) off() string {
	p.isRunning = false
	return "Turning pad off"
}
