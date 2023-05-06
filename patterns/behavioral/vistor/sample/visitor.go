package sample

type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA) string
	VisitConcreteElementB(*ConcreteElementB) int32
}

type Element interface {
	Accept(Visitor)
}

type ConcreteElementA struct {
	Value string
}

func (e *ConcreteElementA) Accept(v Visitor) {
	v.VisitConcreteElementA(e)
}

type ConcreteElementB struct {
	Value int32
}

func (e *ConcreteElementB) Accept(v Visitor) {
	v.VisitConcreteElementB(e)
}

type ConcreteVisitor1 struct{}

func (v *ConcreteVisitor1) VisitConcreteElementA(e *ConcreteElementA) string {
	return e.Value
}

func (v *ConcreteVisitor1) VisitConcreteElementB(e *ConcreteElementB) int32 {
	return e.Value
}

type ConcreteVisitor2 struct{}

func (v *ConcreteVisitor2) VisitConcreteElementA(e *ConcreteElementA) string {
	return e.Value
}

func (v *ConcreteVisitor2) VisitConcreteElementB(e *ConcreteElementB) int32 {
	return e.Value
}
