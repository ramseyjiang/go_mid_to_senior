package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

// SetBuilder allow to change the builder that is being used in the Manufacturing director
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Construct will use the builder that is stored in Manufacturing, and will reproduce the required steps
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
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
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
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
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// Add a new vehicle type is easy.
// type BusBuilder struct {
// 	v VehicleProduct
// }
//
// func (b *BusBuilder) SetWheels() BuildProcess {
// 	b.v.Wheels = 4*2
// 	return b
// }
// func (b *BusBuilder) SetSeats() BuildProcess {
// 	b.v.Seats = 30
// 	return b
// }
// func (b *BusBuilder) SetStructure() BuildProcess {
// 	b.v.Structure = "Bus"
// 	return b
// }
// func (b *BusBuilder) GetVehicle() VehicleProduct {
// 	return b.v
// }
