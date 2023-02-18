package sample

// State is the Step 2. Define an interface for the different states that the object can be in.
type State interface {
	Handle() string
}

// StateA is the Step 3. Implement the different states as concrete types that implement the state interface.
type StateA struct{}

func (s *StateA) Handle() string {
	return "State A"
}

type StateB struct{}

func (s *StateB) Handle() string {
	return "State B"
}

// Content is the Step 1 and 4. Identify the object and add a reference to the state interface.
type Content struct {
	state State
}

// Request is the Step 5. Implement methods in the object that delegate the responsibility of behavior to the state interface.
func (c *Content) Request() string {
	return c.state.Handle()
}

// SetState is the Step 6. Implement methods in the state types that modify the behavior of the object.
func (c *Content) SetState(state State) {
	c.state = state
}
