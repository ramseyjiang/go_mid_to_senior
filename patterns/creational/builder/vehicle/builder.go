package vehicle

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() ProductVehicle
}

type ManufactureDirector struct {
	builder BuildProcess
}

// SetBuilder allow to change the builder that is being used in the Manufacturing director
func (f *ManufactureDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Construct will use the builder that is stored in Manufacturing, and will reproduce the required steps
func (f *ManufactureDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

type ProductVehicle struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v ProductVehicle
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}
func (c *CarBuilder) GetVehicle() ProductVehicle {
	return c.v
}

type BikeBuilder struct {
	v ProductVehicle
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}
func (b *BikeBuilder) GetVehicle() ProductVehicle {
	return b.v
}

// ShuttleBusBuilder is a new vehicle type added is easy.
type ShuttleBusBuilder struct {
	v ProductVehicle
}

func (s *ShuttleBusBuilder) SetWheels() BuildProcess {
	s.v.Wheels = 4 * 2
	return s
}
func (s *ShuttleBusBuilder) SetSeats() BuildProcess {
	s.v.Seats = 30
	return s
}
func (s *ShuttleBusBuilder) SetStructure() BuildProcess {
	s.v.Structure = "ShuttleBus"
	return s
}
func (s *ShuttleBusBuilder) GetVehicle() ProductVehicle {
	return s.v
}
