package area

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestArea(t *testing.T) {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, w: 3}

	areaCalculator := &Calculator{}

	// areaCalculator is the visitor, it is used in calculate() method.
	// The calculate() method can accept many visitors, and it does not need to change the shape interface.
	assert.Equal(t, float32(4), square.calculate(areaCalculator))
	assert.Equal(t, float32(28.26), circle.calculate(areaCalculator))
	assert.Equal(t, float32(6), rectangle.calculate(areaCalculator))
}
