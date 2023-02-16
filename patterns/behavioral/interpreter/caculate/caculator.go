package caculate

import (
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

func getOperationFunc(o string) func(a, b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}
	}

	return nil
}

func isOperator(o string) bool {
	operators := []string{SUM, SUB, DIV, MUL}
	return slices.Contains(operators, o)
}

type polishNotationStack []int

func (p *polishNotationStack) Push(s int) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() int {
	length := len(*p)

	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}

	return 0
}

func Calculate(o string) (int, error) {
	stack := polishNotationStack{}
	operators := strings.Split(o, " ")

	for _, operatorString := range operators {
		if isOperator(operatorString) {
			right := stack.Pop()
			left := stack.Pop()
			// return a func to mathFunc, in other words, mathFunc is a variable to store a func. So mathFunc is used as a func directly.
			mathFunc := getOperationFunc(operatorString)
			res := mathFunc(left, right)
			stack.Push(res)
		} else {
			// The Atoi function takes a string and returns an integer from it or an error.
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				return 0, err
			}

			stack.Push(val)
		}
	}

	return stack.Pop(), nil
}
