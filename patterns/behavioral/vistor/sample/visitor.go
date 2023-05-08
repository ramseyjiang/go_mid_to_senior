package sample

// step 1
type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA) string
	VisitConcreteElementB(*ConcreteElementB) string
}

// step 2
type Element interface {
	Accept(Visitor) string
}

// step 3
type ConcreteElementA struct {
	Value string
}

func (e *ConcreteElementA) Accept(v Visitor) string {
	return v.VisitConcreteElementA(e)
}

type ConcreteElementB struct {
	Value string
}

func (e *ConcreteElementB) Accept(v Visitor) string {
	return v.VisitConcreteElementB(e)
}

// step 4
type ConcreteVisitor1 struct{}

func (v *ConcreteVisitor1) VisitConcreteElementA(e *ConcreteElementA) string {
	return e.Value
}

func (v *ConcreteVisitor1) VisitConcreteElementB(e *ConcreteElementB) string {
	return e.Value
}

type ConcreteVisitor2 struct{}

func (v *ConcreteVisitor2) VisitConcreteElementA(e *ConcreteElementA) string {
	return e.Value
}

func (v *ConcreteVisitor2) VisitConcreteElementB(e *ConcreteElementB) string {
	return e.Value
}
