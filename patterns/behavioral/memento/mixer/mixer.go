package mixer

// UserStory: use a Command pattern to encapsulate a set of different types of states and provide a small facade to automate the insertion
// in the caretaker object.

// Use the same Memento pattern to save two types of states: Volume and Mute.

// Command interface returns the value of its implementer.
// GetValue method returns an interface to a value.
// it returns an interface that can be a representation of any type, and the return should be typecast later if we use it.
type Command interface {
	GetValue() interface{}
}

// Volume state is a byte type
type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

// Mute state a Boolean type
type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

// Memento actor
type Memento struct {
	memento Command
}

// originator actor
type originator struct {
	Command Command
}

func (o *originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

func (o *originator) StoreCommand(m Memento) {
	o.Command = m.memento
}

// careTaker actor
type careTaker struct {
	mementoStack []Memento
}

func (c *careTaker) Push(m Memento) {
	c.mementoStack = append(c.mementoStack, m)
}

func (c *careTaker) Pop() Memento {
	if len(c.mementoStack) > 0 {
		memento := c.mementoStack[len(c.mementoStack)-1]
		c.mementoStack = c.mementoStack[0 : len(c.mementoStack)-1]
		return memento
	}

	return Memento{}
}

// Facade pattern will hold the contents of the originator and the care taker
// and will provide those two easy-to-use methods to save and restore settings

// MementoFacade is using the Facade pattern to automate some tasks
type MementoFacade struct {
	originator originator
	careTaker  careTaker
}

// The SaveSettings method takes a Command, stores it in originator and saves it in an inner careTaker field.
func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Push(m.originator.NewMemento())
}

// The RestoreSettings method makes the opposite flow-restores an index of the careTaker and returns the Command inside the Memento object
func (m *MementoFacade) RestoreSettings() Command {
	m.originator.StoreCommand(m.careTaker.Pop())
	return m.originator.Command
}
