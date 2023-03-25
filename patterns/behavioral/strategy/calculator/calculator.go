package calculator

// Strategy interface
type Strategy interface {
	Execute(a, b int) int
}

// ---------Implement concrete strategies start----------

type AddStrategy struct{}

func (s *AddStrategy) Execute(a, b int) int {
	return a + b
}

type SubtractStrategy struct{}

func (s *SubtractStrategy) Execute(a, b int) int {
	return a - b
}

type MultiplyStrategy struct{}

func (s *MultiplyStrategy) Execute(a, b int) int {
	return a * b
}

type DivideStrategy struct{}

func (s *DivideStrategy) Execute(a, b int) int {
	return a / b
}

// ---------Implement concrete strategies end----------

// ---------Create a context to use the strategy start

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

// ---------Create a context to use the strategy end
