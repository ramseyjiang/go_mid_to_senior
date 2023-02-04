package sample

// Prototype interface
type Prototype interface {
	Clone() Prototype
	GetName() string
}

type ConcretePrototypeA struct {
	Name string
}

func (a *ConcretePrototypeA) Clone() Prototype {
	return &ConcretePrototypeA{Name: a.Name}
}

func (a *ConcretePrototypeA) GetName() string {
	return a.Name
}

type ConcretePrototypeB struct {
	Name string
}

func (b *ConcretePrototypeB) Clone() Prototype {
	return &ConcretePrototypeB{Name: b.Name}
}

func (b *ConcretePrototypeB) GetName() string {
	return b.Name
}
