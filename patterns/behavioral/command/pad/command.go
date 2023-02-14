package pad

// Command interface
type Command interface {
	execute() string
}

// OnCommand is a Concrete command
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() string {
	return c.device.on()
}

// OffCommand is a Concrete command
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() string {
	return c.device.off()
}
