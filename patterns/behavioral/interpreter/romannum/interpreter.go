package romannum

import "strings"

type Expression interface {
	Interpret(context string) bool
}

type TerminalExpression struct {
	Data string
}

func (te *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, te.Data) {
		return true
	}
	return false
}

type OrExpression struct {
	Expr1 Expression
	Expr2 Expression
}

func (oe *OrExpression) Interpret(context string) bool {
	return oe.Expr1.Interpret(context) || oe.Expr2.Interpret(context)
}
