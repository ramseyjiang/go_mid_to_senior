package transportation

// LuxuryCar is Concrete product
type LuxuryCar struct{}

func (*LuxuryCar) NumDoors() int {
	return 4
}
func (*LuxuryCar) NumWheels() int {
	return 4
}
func (*LuxuryCar) NumSeats() int {
	return 5
}
