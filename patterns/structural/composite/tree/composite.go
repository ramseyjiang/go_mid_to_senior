package tree

import (
	"log"
)

// Component interface, which has a single method Operation(). Leaf and Composite are two types of components.
type Component interface {
	Operation()
}

// Leaf represents an individual object
type Leaf struct {
}

func (l *Leaf) Operation() {
	log.Println("Leaf operation")
}

// Composite represents a group of objects. Composite type has a children field that stores a slice of components.
// It has three methods: Add(), Remove(), and Operation().
// The Add() and Remove() methods allow you to add and remove components from the composite, respectively.
type Composite struct {
	children []Component
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Remove(component Component) {
	for i, child := range c.children {
		if child == component {
			c.children = append(c.children[:i], c.children[i+1:]...)
			break
		}

		if composite, ok := child.(*Composite); ok {
			composite.Remove(child)
		}
	}
}

// Operation method performs an operation on the composite and its children.
func (c *Composite) Operation() {
	log.Println("Composite operation")
	for _, child := range c.children {
		child.Operation()
	}
}
