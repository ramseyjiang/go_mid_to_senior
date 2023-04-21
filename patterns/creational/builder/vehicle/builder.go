package vehicle

// Step 1: Define the Product with the required properties.

type Product struct {
	Wheels    int
	Seats     int
	Structure string
}

// Step 2: Create a Builder interface with methods to construct the object

type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	GetVehicle() Product
}

// CarBuilder is used to concrete Builder structs, it is step 3
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

func (c *CarBuilder) GetVehicle() Product {
	return c.v
}

// BikeBuilder is used to concrete Builder structs
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

// ShuttleBusBuilder is used to concrete Builder structs
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

// ManufactureDirector is used to take a Builder interface and construct the object.
type ManufactureDirector struct {
	builder Builder
}

func (f *ManufactureDirector) SetBuilder(b Builder) {
	f.builder = b
}

func (f *ManufactureDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}
