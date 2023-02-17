package area

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestArea(t *testing.T) {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, w: 3}

	areaCalculator := &CalculatorArea{}

	// areaCalculator is the visitor, it is used in calculate() method.
	// The calculate() method can accept many visitors, and it does not need to change the shape interface.
	assert.Equal(t, float32(4), square.calculate(areaCalculator))
	assert.Equal(t, float32(28.26), circle.calculate(areaCalculator))
	assert.Equal(t, float32(6), rectangle.calculate(areaCalculator))

	// perimeterCalculator is another visitor, it is used in calculate() method.
	// Use different visitors can do different things, and don't need to update the source code in the shape.
	perimeterCalculator := &CalculatorPerimeter{}
	assert.Equal(t, float32(8), square.calculate(perimeterCalculator))
	assert.Equal(t, float32(18.84), circle.calculate(perimeterCalculator))
	assert.Equal(t, float32(10), rectangle.calculate(perimeterCalculator))
}
