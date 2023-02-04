package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPrototypeClone(t *testing.T) {
	prototypeA := &ConcretePrototypeA{Name: "Prototype A"}
	prototypeB := &ConcretePrototypeB{Name: "Prototype B"}

	clonedPrototypeA := prototypeA.Clone()
	clonedPrototypeB := prototypeB.Clone()

	assert.Equal(t, "Prototype A", clonedPrototypeA.GetName())
	assert.Equal(t, "Prototype B", clonedPrototypeB.GetName())
}
