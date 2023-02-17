package area

// CalculatorArea is one of concrete visitor, the below methods implement the Visitor abstract interface.
type CalculatorArea struct {
	area float32
}

func (c *CalculatorArea) visitSquare(s *Square) float32 {
	return s.side * s.side
}

func (c *CalculatorArea) visitCircle(s *Circle) float32 {
	return 3.14 * s.radius * s.radius
}

func (c *CalculatorArea) visitRectangle(s *Rectangle) float32 {
	return s.l * s.w
}

// CalculatorPerimeter is another concrete visitor. The below methods implement the Visitor abstract interface.
type CalculatorPerimeter struct {
	perimeter float32
}

func (p *CalculatorPerimeter) visitSquare(s *Square) float32 {
	return s.side * 4
}

func (p *CalculatorPerimeter) visitCircle(s *Circle) float32 {
	return 3.14 * s.radius * 2
}

func (p *CalculatorPerimeter) visitRectangle(s *Rectangle) float32 {
	return s.l*2 + s.w*2
}
