package factory

import (
	"fmt"
)

const Bucks = "Bucks"
const Slipper = "Slipper"
const Sandal = "Sandal"

func getFootWear(category string, size int64, gender string, price float32) (iFootwear, error) {
	switch category {
	case Slipper:
		return newShoe(size, Slipper, price, gender, 0.5), nil
	case Sandal:
		return newShoe(size, Sandal, price, gender, 0.7), nil
	case Bucks:
		return newShoe(size, Bucks, price, gender, 0), nil
	default:
		return nil, fmt.Errorf("invalid footwear type")
	}
}

func Entry() {
	footwear1, _ := getFootWear("Slipper", 41, "Male", 200)
	fmt.Println(footwear1.getPrice(), footwear1.getCategory(), footwear1.getGender())

	footwear2, _ := getFootWear("Sandal", 42, "Male", 1000)
	fmt.Println(footwear2.getPrice(), footwear2.getCategory(), footwear2.getGender())

	footwear3, _ := getFootWear("Bucks", 38, "Female", 100)
	fmt.Println(footwear3.getPrice(), footwear3.getCategory(), footwear3.getGender())

	footwear4, err := getFootWear("test", 38, "Female", 100)
	fmt.Println(footwear4, err)
}
