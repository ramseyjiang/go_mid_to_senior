package cmd

type device interface {
	on()
	off()
	increaseVolume()
	decreaseVolume()
}

type command interface {
	execute()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type increaseVolumeCommand struct {
	device device
}

func (c *increaseVolumeCommand) execute() {
	c.device.increaseVolume()
}

type decreaseVolumeCommand struct {
	device device
}

func (c *decreaseVolumeCommand) execute() {
	c.device.decreaseVolume()
}

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// Trigger go run patterns/command/*.go
func Trigger() {
	tv := &tv{
		isOn:   false,
		volume: 10,
	}

	// Instantiate commands
	onCommand := &onCommand{
		device: tv,
	}
	offCommand := &offCommand{
		device: tv,
	}
	increaseVolumeCommand := &increaseVolumeCommand{
		device: tv,
	}
	decreaseVolumeCommand := &decreaseVolumeCommand{
		device: tv,
	}

	// Instantiate buttons
	onButton := &button{
		command: onCommand,
	}
	offButton := &button{
		command: offCommand,
	}
	increaseVolumeButton := &button{
		command: increaseVolumeCommand,
	}
	decreaseVolumeButton := &button{
		command: decreaseVolumeCommand,
	}

	// Execute
	increaseVolumeButton.press()
	onButton.press()
	increaseVolumeButton.press()
	decreaseVolumeButton.press()
	offButton.press()
}
