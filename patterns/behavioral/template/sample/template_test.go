package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// create instances of the two concrete classes and call the TemplateMethod() function to execute the algorithm.
func TestTemplate(t *testing.T) {
	// Call template method for concrete class 1
	c1 := &ConcreteClass1{}
	r1 := c1.TemplateMethod()
	assert.Equal(t, r1[0], "Concrete class 1 - Step 1 done")
	assert.Equal(t, r1[1], "Concrete class 1 - Step 2 done")

	// Call template method for concrete class 2
	c2 := &ConcreteClass2{}
	r2 := c2.TemplateMethod()
	assert.Equal(t, r2[0], "Concrete class 2 - Step 1 done")
	assert.Equal(t, r2[1], "Concrete class 2 - Step 2 done")
	assert.Equal(t, r2[2], "Concrete class 2 - Step 3 done")
}
