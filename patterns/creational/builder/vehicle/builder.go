package vehicle

// Builder is the interface that specifies the methods for building the parts of the product
type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	GetVehicle() Product
}

// Product represents the complex object that we want to build
type Product struct {
	Wheels    int
	Seats     int
	Structure string
}

// ManufactureDirector is responsible for using the Builder interface to build the product
type ManufactureDirector struct {
	builder Builder
}

// SetBuilder allow to change the builder that is being used in the Manufacturing director
func (f *ManufactureDirector) SetBuilder(b Builder) {
	f.builder = b
}

// Construct will use the builder that is stored in Manufacturing, and will reproduce the required steps
func (f *ManufactureDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// CarBuilder is an implementation of the Builder interface
type CarBuilder struct {
	v Product
}

func (c *CarBuilder) SetWheels() Builder {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() Builder {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() Builder {
	c.v.Structure = "Car"
	return c
}

// GetVehicle returns the built product
func (c *CarBuilder) GetVehicle() Product {
	return c.v
}

// BikeBuilder is an implementation of the Builder interface
type BikeBuilder struct {
	v Product
}

func (b *BikeBuilder) SetWheels() Builder {
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() Builder {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) SetStructure() Builder {
	b.v.Structure = "Bike"
	return b
}
func (b *BikeBuilder) GetVehicle() Product {
	return b.v
}

// ShuttleBusBuilder is an implementation of the Builder interface
type ShuttleBusBuilder struct {
	v Product
}

func (s *ShuttleBusBuilder) SetWheels() Builder {
	s.v.Wheels = 4 * 2
	return s
}
func (s *ShuttleBusBuilder) SetSeats() Builder {
	s.v.Seats = 30
	return s
}
func (s *ShuttleBusBuilder) SetStructure() Builder {
	s.v.Structure = "ShuttleBus"
	return s
}
func (s *ShuttleBusBuilder) GetVehicle() Product {
	return s.v
}
