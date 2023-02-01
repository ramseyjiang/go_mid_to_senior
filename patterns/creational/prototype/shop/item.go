package shop

import (
	"fmt"

	"github.com/google/uuid"
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SN    string
	Color ShirtColor
}

func (i *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f", i.SN, i.Color, i.Price)
}

var whitePrototype = &Shirt{
	Price: 15.00,
	SN:    uuid.New().String(),
	Color: White,
}

var blackPrototype = &Shirt{
	Price: 16.00,
	SN:    uuid.New().String(),
	Color: Black,
}

var bluePrototype = &Shirt{
	Price: 17.00,
	SN:    uuid.New().String(),
	Color: Blue,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}
