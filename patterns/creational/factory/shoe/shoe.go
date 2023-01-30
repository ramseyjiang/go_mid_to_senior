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

func getShoe(category string, size int64, gender string, price float32) (footwear, error) {
	switch category {
	case Slipper:
		return newShoe(size, Slipper, price, gender, SlipperDiscount), nil
	case Sandal:
		return newShoe(size, Sandal, price, gender, SandalDiscount), nil
	case Bucks:
		return newShoe(size, Bucks, price, gender, BucksDiscount), nil
	default:
		return nil, fmt.Errorf("invalid footwear type")
	}
}
