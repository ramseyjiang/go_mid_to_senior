package noodle

type Noodle interface {
	getPrice() int
}

type VegetableMania struct {
}

func (v *VegetableMania) getPrice() int {
	return 15
}

type CheeseTopping struct {
	noodle Noodle
}

func (c *CheeseTopping) getPrice() int {
	noodlePrice := c.noodle.getPrice()
	return noodlePrice + 7
}

type TomatoTopping struct {
	noodle Noodle
}

func (c *TomatoTopping) getPrice() int {
	noodlePrice := c.noodle.getPrice()
	return noodlePrice + 10
}
