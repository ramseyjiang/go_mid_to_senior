package area

// Shape interface defines. calculate(Visitor) is the key method in the visitor pattern.
// It can accept many visitors and does not need to change the Shape interface.
type Shape interface {
	getType() string
	calculate(Visitor) float32
}

// Visitor is the abstract interface
type Visitor interface {
	visitSquare(*Square) float32
	visitCircle(*Circle) float32
	visitRectangle(*Rectangle) float32
}

// Square is the concrete element of the shape
type Square struct {
	side float32
}

func (s *Square) calculate(v Visitor) float32 {
	return v.visitSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

// Circle is the concrete element of the shape
type Circle struct {
	radius float32
}

func (c *Circle) calculate(v Visitor) float32 {
	return v.visitCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

// Rectangle is the concrete element of the shape
type Rectangle struct {
	l float32
	w float32
}

func (t *Rectangle) calculate(v Visitor) float32 {
	return v.visitRectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}
