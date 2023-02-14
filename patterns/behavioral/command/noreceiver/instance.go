package noreceiver

// This instance does not have an receiver interface.

// Command interface
type Command interface {
	Execute() string
}

// ConsoleOutput is a concrete command
type ConsoleOutput struct {
	message string
}

// Execute is the implement of the Command interface
func (c *ConsoleOutput) Execute() string {
	return c.message
}

func CreateCommand(s string) Command {
	return &ConsoleOutput{
		message: s,
	}
}

// CommandQueue is the invoker struct
type CommandQueue struct {
	queue  []Command
	record []string
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 {
		for _, command := range p.queue {
			p.record = append(p.record, command.Execute())
		}

		p.queue = make([]Command, 3)
	}
}
