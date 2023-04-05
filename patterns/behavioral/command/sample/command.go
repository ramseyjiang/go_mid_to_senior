package sample

type Command interface {
	Execute() string
}

type ConcreteCommandA struct {
	receiver *Receiver
}

func (c *ConcreteCommandA) Execute() string {
	return c.receiver.ActionA()
}

type ConcreteCommandB struct {
	receiver *Receiver
}

func (c *ConcreteCommandB) Execute() string {
	return c.receiver.ActionB()
}

type Receiver struct{}

func (r *Receiver) ActionA() string {
	return "ActionA called"
}

func (r *Receiver) ActionB() string {
	return "ActionB called"
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(c Command) {
	i.command = c
}

func (i *Invoker) ExecuteCommand() string {
	return i.command.Execute()
}
