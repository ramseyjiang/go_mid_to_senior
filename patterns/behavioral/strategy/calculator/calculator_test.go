package calculator

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCalculator(t *testing.T) {
	// Use the context and strategies
	addStrategy := &AddStrategy{}
	subtractStrategy := &SubtractStrategy{}
	multiplyStrategy := &MultiplyStrategy{}
	divideStrategy := &DivideStrategy{}

	// Use addition strategy
	context := NewContext(addStrategy)
	assert.Equal(t, 15, context.ExecuteStrategy(10, 5))

	// Use subtraction strategy
	context = NewContext(subtractStrategy)
	assert.Equal(t, 5, context.ExecuteStrategy(10, 5))

	// Use multiplication strategy
	context = NewContext(multiplyStrategy)
	assert.Equal(t, 50, context.ExecuteStrategy(10, 5))

	// Use division strategy
	context = NewContext(divideStrategy)
	assert.Equal(t, 2, context.ExecuteStrategy(10, 5))
}
