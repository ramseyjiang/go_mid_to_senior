package pad

type Button struct {
	command Command
}

func (b *Button) press() string {
	return b.command.execute()
}
