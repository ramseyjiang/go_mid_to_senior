package caculator

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// create instances of the Number class for the values 2, 3, and 4,
// and then use them to build a parse tree for the expression 2 + 3 - 4.
// We can then use the Interpret method to evaluate the expression and produce a result of 1.
func TestExpression(t *testing.T) {
	two := &Number{2}
	three := &Number{3}
	four := &Number{4}

	result := (&Minus{
		left: &Plus{
			left:  two,
			right: three,
		},
		right: four,
	}).Interpret()

	assert.Equal(t, 1, result)
}
