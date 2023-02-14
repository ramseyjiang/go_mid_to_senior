package sample

type Command interface {
	Execute() string
}

type ConcreteCommand struct {
	receiver Receiver
}

func (c *ConcreteCommand) Execute() string {
	return c.receiver.Action()
}

type Receiver interface {
	Action() string
}

type ConcreteReceiver struct{}

func (r *ConcreteReceiver) Action() string {
	return "Action called"
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
