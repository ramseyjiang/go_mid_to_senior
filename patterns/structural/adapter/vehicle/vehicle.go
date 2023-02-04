package vehicle

// Vehicle is the interface that Client uses
type Vehicle interface {
	Drive() string
}

// Car is an existing struct with a specific interface
type Car struct {
	Name string
}

func (c *Car) RunOnRoad() string {
	return c.Name + " is driving on road."
}

// Plane is an existing struct with a specific interface
type Plane struct {
	Name string
}

func (p *Plane) FlyInSky() string {
	return p.Name + " is flying in sky."
}

// Boat is an existing struct with a specific interface
type Boat struct {
	Name string
}

func (b *Boat) SailOnWater() string {
	return b.Name + " is sailing on water."
}

// CarAdapter allows the Car to be used as a Vehicle
type CarAdapter struct {
	car *Car
}

func (c *CarAdapter) Drive() string {
	return c.car.RunOnRoad()
}

// PlaneAdapter allows the Plane to be used as a Vehicle
type PlaneAdapter struct {
	plane *Plane
}

func (p *PlaneAdapter) Drive() string {
	return p.plane.FlyInSky()
}

// BoatAdapter allows the Boat to be used as a Vehicle
type BoatAdapter struct {
	boat *Boat
}

func (b *BoatAdapter) Drive() string {
	return b.boat.SailOnWater()
}
