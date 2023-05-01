package calculator

type Expression interface {
	Interpret() int
}

// Number is an expression struct, it represents a numeric value struct,
type Number struct {
	value int
}

func (n *Number) Interpret() int {
	return n.value
}

// Plus is an expression struct, Plus and Minus represent addition and subtraction, respectively.
type Plus struct {
	left  Expression
	right Expression
}

func (p *Plus) Interpret() int {
	return p.left.Interpret() + p.right.Interpret()
}

// Minus is an expression struct
type Minus struct {
	left  Expression
	right Expression
}

func (m *Minus) Interpret() int {
	return m.left.Interpret() - m.right.Interpret()
}
