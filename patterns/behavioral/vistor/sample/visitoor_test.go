package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestVisitor(t *testing.T) {
	var value1 = "Test"
	var value2 int32 = 101
	visitor1 := &ConcreteVisitor1{}
	visitor2 := &ConcreteVisitor2{}

	element1 := &ConcreteElementA{Value: "Test"}
	element2 := &ConcreteElementB{Value: 101}

	t.Run("Test VisitConcreteElementA with ConcreteVisitor1", func(t *testing.T) {
		assert.Equal(t, visitor1.VisitConcreteElementA(element1), value1)
	})

	t.Run("Test VisitConcreteElementB with ConcreteVisitor1", func(t *testing.T) {
		assert.Equal(t, visitor1.VisitConcreteElementB(element2), value2)
	})

	t.Run("Test VisitConcreteElementA with ConcreteVisitor2", func(t *testing.T) {
		assert.Equal(t, visitor2.VisitConcreteElementA(element1), value1)
	})

	t.Run("Test VisitConcreteElementB with ConcreteVisitor2", func(t *testing.T) {
		assert.Equal(t, visitor2.VisitConcreteElementB(element2), value2)
	})
}
