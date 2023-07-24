package shoe

import (
	"fmt"
)

const Bucks = "Bucks"
const Slipper = "Slipper"
const Sandal = "Sandal"

const SlipperDiscount = 0.5
const SandalDiscount = 0.7
const BucksDiscount = 0

type FactoryShoe interface {
	CreateShoe(shoeType string, size int64, gender string, price float32) (Footwear, error)
}

type ConcreteShoeFactory struct{}

func (c *ConcreteShoeFactory) CreateShoe(category string, size int64, gender string, price float32) (Footwear, error) {
	switch category {
	case Slipper:
		return NewShoe(Slipper, size, price, gender, SlipperDiscount), nil
	case Sandal:
		return NewShoe(Sandal, size, price, gender, SandalDiscount), nil
	case Bucks:
		return NewShoe(Bucks, size, price, gender, BucksDiscount), nil
	default:
		return nil, fmt.Errorf("invalid footwear type")
	}
}

func NewConcreteShoeFactory() FactoryShoe {
	return &ConcreteShoeFactory{}
}
