package delivery

type State interface {
	Handle() string
}

type Package struct {
	state State
}

func NewPackage() *Package {
	return &Package{
		state: &OrderedState{},
	}
}

func (p *Package) UpdateState(s State) {
	p.state = s
}

func (p *Package) CurrentState() string {
	return p.state.Handle()
}

type OrderedState struct{}

func (o *OrderedState) Handle() string {
	return "Package is ordered and is waiting for courier."
}

type ShippedState struct{}

func (s *ShippedState) Handle() string {
	return "Package is shipped and is in transit."
}

type DeliveredState struct{}

func (d *DeliveredState) Handle() string {
	return "Package is delivered to the recipient."
}
