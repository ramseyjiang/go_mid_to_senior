package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestVisitor(t *testing.T) {
	var value1 = "Test1"
	var value2 = "Test2"

	elementA := &ConcreteElementA{Value: value1}
	elementB := &ConcreteElementB{Value: value2}

	visitor1 := &ConcreteVisitor1{}
	visitor2 := &ConcreteVisitor2{}

	expectedValue1 := "Test1"
	expectedValue2 := "Test2"

	t.Run("Test VisitConcreteElementA with ConcreteVisitor1", func(t *testing.T) {
		assert.Equal(t, elementA.Accept(visitor1), expectedValue1)
	})

	t.Run("Test VisitConcreteElementB with ConcreteVisitor1", func(t *testing.T) {
		assert.Equal(t, elementB.Accept(visitor1), expectedValue2)
	})

	t.Run("Test VisitConcreteElementA with ConcreteVisitor2", func(t *testing.T) {
		assert.Equal(t, elementA.Accept(visitor2), expectedValue1)
	})

	t.Run("Test VisitConcreteElementB with ConcreteVisitor2", func(t *testing.T) {
		assert.Equal(t, elementB.Accept(visitor2), expectedValue2)
	})
}
