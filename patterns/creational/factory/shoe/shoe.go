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
	CreateShoe(shoeType string) (Footwear, error)
}

type ConcreteShoeFactory struct{}

func CreateShoe(category string, size int64, gender string, price float32) (Footwear, error) {
	switch category {
	case Slipper:
		return newShoe(Slipper, size, price, gender, SlipperDiscount), nil
	case Sandal:
		return newShoe(Sandal, size, price, gender, SandalDiscount), nil
	case Bucks:
		return newShoe(Bucks, size, price, gender, BucksDiscount), nil
	default:
		return nil, fmt.Errorf("invalid footwear type")
	}
}
