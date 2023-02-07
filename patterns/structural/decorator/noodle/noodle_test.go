package noodle

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNoodle(t *testing.T) {
	noodle := &VegetableMania{}

	// Add fried Noodle
	noodleWithCheese := &CheeseTopping{
		noodle: noodle,
	}

	// Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		noodle: noodleWithCheese,
	}

	assert.Equal(t, 32, pizzaWithCheeseAndTomato.getPrice())
}
