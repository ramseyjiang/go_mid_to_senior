package sample

// In this implementation, the AbstractClass is defined as an interface with a TemplateMethod() function that defines the steps of the algorithm.
// Two concrete classes, ConcreteClass1 and ConcreteClass2, are defined and implement the TemplateMethod() function with specific behaviors that can vary.

// AbstractClass Abstract class
type AbstractClass interface {
	TemplateMethod()
}

// ConcreteClass1 Concrete class 1
type ConcreteClass1 struct {
	record []string
}

func (c1 *ConcreteClass1) TemplateMethod() []string {
	c1.record = append(c1.record, "Concrete class 1 - Step 1 done")
	c1.record = append(c1.record, "Concrete class 1 - Step 2 done")

	return c1.record
}

// ConcreteClass2 Concrete class 2
type ConcreteClass2 struct {
	record []string
}

func (c2 *ConcreteClass2) TemplateMethod() []string {
	c2.record = append(c2.record, "Concrete class 2 - Step 1 done")
	c2.record = append(c2.record, "Concrete class 2 - Step 2 done")
	c2.record = append(c2.record, "Concrete class 2 - Step 3 done")

	return c2.record
}
