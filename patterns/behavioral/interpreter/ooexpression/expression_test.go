package ooexpression

import (
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestExpression(t *testing.T) {
	stack := polishNotationStack{}
	operators := strings.Split("3 4 sum 2 sub", " ")

	for _, operatorString := range operators {
		if operatorString == SUM || operatorString == SUB {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := operatorFactory(operatorString, left, right)
			res := value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				panic(err)
			}

			temp := value(val)
			stack.Push(&temp)
		}
	}

	assert.Equal(t, 5, int(stack.Pop().Read()))
}
