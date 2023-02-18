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

	// areaCalculator is the visitor, it is used in accept() method.
	// The accept() method can accept many visitors, and it does not need to change the shape interface.
	assert.Equal(t, float32(4), square.Accept(areaCalculator))
	assert.Equal(t, float32(28.26), circle.Accept(areaCalculator))
	assert.Equal(t, float32(6), rectangle.Accept(areaCalculator))

	// perimeterCalculator is another visitor, it is used in accept() method.
	// Use different visitors can do different things, and don't need to update the source code in the shape.
	perimeterCalculator := &CalculatorPerimeter{}
	assert.Equal(t, float32(8), square.Accept(perimeterCalculator))
	assert.Equal(t, float32(18.84), circle.Accept(perimeterCalculator))
	assert.Equal(t, float32(10), rectangle.Accept(perimeterCalculator))
}
