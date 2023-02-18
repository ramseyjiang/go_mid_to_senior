package shop

import (
	"testing"

	"golang.org/x/exp/slices"

	"github.com/go-playground/assert/v2"
)

func TestShop(t *testing.T) {
	products := make([]Visitable, 3)
	products[0] = &Rice{
		Product: Product{
			Price: 22.0,
			Name:  "Some rice",
		},
	}
	products[1] = &Pasta{
		Product: Product{
			Price: 50.0,
			Name:  "Some pasta",
		},
	}
	products[2] = &Fridge{
		Product: Product{
			Price: 70,
			Name:  "A fridge",
		},
	}

	priceVisitor := &PriceVisitor{}

	for _, p := range products {
		p.Accept(priceVisitor)
	}

	assert.Equal(t, float32(172), priceVisitor.Sum)

	nameVisitor := &NameVisitor{}

	for _, p := range products {
		p.Accept(nameVisitor)
	}

	assert.Equal(t, true, slices.Contains(nameVisitor.ProductList, "A fridge"))
	assert.Equal(t, true, slices.Contains(nameVisitor.ProductList, "Some pasta"))
	assert.Equal(t, true, slices.Contains(nameVisitor.ProductList, "Some rice"))
}
