package shapes

type Geometry interface {
	Area() float64
	Perimeter() float64
	Width() float64
	Height() float64
	SetWidth(float64) Geometry
	SetHeight(float64) Geometry
}

type Square struct {
	side float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (s *Square) Perimeter() float64 {
	return s.side * 4
}

func (s *Square) Width() float64 {
	return s.side
}

func (s *Square) Height() float64 {
	return s.side
}

func (s *Square) SetWidth(f float64) Geometry {
	s.side = f
	return s
}

func (s *Square) SetHeight(f float64) Geometry {
	s.side = f
	return s
}

type Rectangle struct {
	width  float64
	height float64
}

func (s *Rectangle) Area() float64 {
	return s.width * s.height
}

func (s *Rectangle) Perimeter() float64 {
	return 2 * (s.width + s.height)
}

func (s *Rectangle) Width() float64 {
	return s.width
}

func (s *Rectangle) Height() float64 {
	return s.height
}

func (s *Rectangle) SetWidth(f float64) Geometry {
	s.width = f
	return s
}

func (s *Rectangle) SetHeight(f float64) Geometry {
	s.height = f
	return s
}
