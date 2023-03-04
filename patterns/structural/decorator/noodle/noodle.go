package noodle

// Noodle is an interface that specifies the methods that all decorators will implement.
type Noodle interface {
	getPrice() int
}

// VegetableMania is a struct that represents the base object that will be decorated.
type VegetableMania struct {
}

// Implement the methods for the base object.
func (v *VegetableMania) getPrice() int {
	return 15
}

type Bonus interface {
	Noodle
	BuyTwoBonusOne() int
}

// CheeseTopping is the struct represents a decorator
type CheeseTopping struct {
	noodle Noodle
	price  int
}

// getPrice implement the method for the decorator struct.
func (c *CheeseTopping) getPrice() int {
	c.price = c.noodle.getPrice() + 7
	return c.price
}

func (c *CheeseTopping) BuyTwoBonusOne() int {
	return c.price * 2
}

// TomatoTopping is the struct represents a decorator
type TomatoTopping struct {
	noodle Noodle
	price  int
}

// Embed the base object in the decorator, implement the methods for the decorator struct.
func (t *TomatoTopping) getPrice() int {
	t.price = t.noodle.getPrice() + 10
	return t.price
}

func (t *TomatoTopping) BuyTwoBonusOne() int {
	return t.price * 2
}
