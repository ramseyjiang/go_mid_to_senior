package noodle

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNoodle(t *testing.T) {
	noodle := &VegetableMania{}

	noodleWithCheese := &CheeseTopping{
		noodle: noodle,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		noodle: noodleWithCheese,
	}

	assert.Equal(t, 15, noodle.getPrice())
	assert.Equal(t, 22, noodleWithCheese.getPrice())
	assert.Equal(t, 32, pizzaWithCheeseAndTomato.getPrice())
}
