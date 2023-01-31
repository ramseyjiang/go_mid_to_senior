package clothes

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdidasShoe(t *testing.T) {
	adidasFactory, _ := GetSportsFactory(adidas)
	adidasShoe := adidasFactory.makeShoe()
	assert.Equal(t, adidas, adidasShoe.getLogo())
	assert.Equal(t, 14, adidasShoe.getSize())
}

func TestAdidasShirt(t *testing.T) {
	adidasFactory, _ := GetSportsFactory(adidas)
	adidasShirt := adidasFactory.makeShirt()
	assert.Equal(t, adidas, adidasShirt.getLogo())
	assert.Equal(t, 14, adidasShirt.getSize())
}

func TestNikeShoe(t *testing.T) {
	nikeFactory, _ := GetSportsFactory(nike)
	nikeShoe := nikeFactory.makeShoe()
	assert.Equal(t, nike, nikeShoe.getLogo())
	assert.Equal(t, 14, nikeShoe.getSize())
}

func TestNikeShirt(t *testing.T) {
	nikeFactory, _ := GetSportsFactory(nike)
	nikeShirt := nikeFactory.makeShirt()
	assert.Equal(t, nike, nikeShirt.getLogo())
	assert.Equal(t, 14, nikeShirt.getSize())
}
