package icecream

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetIceCream(t *testing.T) {
	appleCream := GetIceCream("apple")
	berryCream := GetIceCream("strawberry")
	watermelonCream := GetIceCream("watermelon")

	assert.Equal(t, "Favour is apple", appleCream.Favour)
	assert.Equal(t, "Favour is strawberry", berryCream.Favour)
	assert.Equal(t, "Favour is watermelon", watermelonCream.Favour)
}
