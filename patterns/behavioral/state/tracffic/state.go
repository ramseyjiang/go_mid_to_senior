package tracffic

// State is the first step to define the State interface
type State interface {
	Handle(context *Context) string
}

// Context is the second step to define the Context Struct.
type Context struct {
	state State
}

func NewContext(state State) *Context {
	return &Context{
		state: state,
	}
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() string {
	return c.state.Handle(c)
}

// GreenLightState YellowLightState, RedLightState are all concretes structs to implement the state interface.
type GreenLightState struct{}

func (g *GreenLightState) Handle(context *Context) string {
	context.SetState(new(RedLightState))
	return "Green Light - Go!"
}

type YellowLightState struct{}

func (y *YellowLightState) Handle(context *Context) string {
	context.SetState(new(GreenLightState))
	return "Yellow Light - Wait!"
}

type RedLightState struct{}

func (r *RedLightState) Handle(context *Context) string {
	context.SetState(new(YellowLightState))
	return "Red Light - Stop!"
}
