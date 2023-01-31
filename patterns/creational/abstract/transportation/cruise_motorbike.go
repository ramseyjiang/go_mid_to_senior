package transportation

// CruiseMotorbike is Concrete product
type CruiseMotorbike struct{}

func (c *CruiseMotorbike) NumWheels() int {
	return 2
}
func (c *CruiseMotorbike) NumSeats() int {
	return 2
}
func (c *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
