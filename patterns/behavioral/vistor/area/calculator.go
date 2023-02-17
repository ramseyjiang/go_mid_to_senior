package area

// Calculator is one of concrete struct, the below methods implement the Visitor abstract interface.
type Calculator struct {
	area float32
}

func (c *Calculator) visitSquare(s *Square) float32 {
	return s.side * s.side
}

func (c *Calculator) visitCircle(s *Circle) float32 {
	return 3.14 * s.radius * s.radius
}

func (c *Calculator) visitRectangle(s *Rectangle) float32 {
	return s.l * s.w
}
