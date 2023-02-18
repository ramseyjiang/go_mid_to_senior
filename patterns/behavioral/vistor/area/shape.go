package area

// Shape interface defines.
type Shape interface {
	getType() string
	// Accept(Visitor) Someone likes to define Accept(Visitor) at here, you should update the original interface, it's not a good behaviour.
}

// The Visitable interface has a method called Accept(Visitor) that will execute the decoupled algorithm.
// Accept(Visitor) is the key method in the visitor pattern. It is always defined using Accept().
// Visitor is the abstract interface, all visit methods are defined in it.
// Methods in Visitor interface can be used to alter the Visitable object.
// Using this way, Visitor can change the original functionalities indirectly, and it won't change the original interface source code.
type Visitable interface {
	Accept(Visitor) float32
}

// Visitor is the abstract interface
type Visitor interface {
	visitSquare(*Square) float32
	visitCircle(*Circle) float32
	visitRectangle(*Rectangle) float32
}

// Square is the concrete element of the shape, it will implement two methods which from Shape interface and Visitable interface.
type Square struct {
	side float32
}

func (s *Square) Accept(v Visitor) float32 {
	return v.visitSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

// Circle is the concrete element of the shape, it will implement two methods which from Shape interface and Visitable interface.
type Circle struct {
	radius float32
}

func (c *Circle) Accept(v Visitor) float32 {
	return v.visitCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

// Rectangle is the concrete element of the shape, it will implement two methods which from Shape interface and Visitable interface.
type Rectangle struct {
	l float32
	w float32
}

func (t *Rectangle) Accept(v Visitor) float32 {
	return v.visitRectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}
