package sample

// The Memento struct stores the state of the Originator object.
type Memento struct {
	state string
}

// The Originator struct is the object whose state we want to save and restore
// If the Memento contains a specific state, the originator type contains the state that is currently loaded.
// Also, to save the state of something could be as simple as to take some value or as complex as to maintain the state of some distributed application.
type Originator struct {
	state string
}

func (o *Originator) setState(state string) {
	o.state = state
}

func (o *Originator) getState() string {
	return o.state
}

// createMemento returns a new Memento object containing the current state.
func (o *Originator) createMemento() *Memento {
	return &Memento{state: o.state}
}

// restoreMemento takes a Memento object and sets the Originator's state to the state stored in the Memento.
func (o *Originator) restoreMemento(m *Memento) {
	o.state = m.state
}

// CareTaker struct, which is responsible for storing and retrieving Memento objects
type CareTaker struct {
	mementos []*Memento
}

func (c *CareTaker) addMemento(m *Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *CareTaker) getMemento(index int) *Memento {
	if len(c.mementos) < index || index < 0 {
		return &Memento{}
	}
	return c.mementos[index]
}
