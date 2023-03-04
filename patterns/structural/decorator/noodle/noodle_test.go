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

	noodleWithCheeseAndTomato := &TomatoTopping{
		noodle: noodleWithCheese,
	}

	assert.Equal(t, 15, noodle.getPrice())
	assert.Equal(t, 22, noodleWithCheese.getPrice())
	assert.Equal(t, 44, noodleWithCheese.BuyTwoBonusOne())
	assert.Equal(t, 32, noodleWithCheeseAndTomato.getPrice())
	assert.Equal(t, 64, noodleWithCheeseAndTomato.BuyTwoBonusOne())
}
