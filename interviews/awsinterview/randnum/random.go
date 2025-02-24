package main

import (
	"fmt"
	"math/rand"
)

func main() {
	operators := []rune{'+', '-', '*', '/'}
	op := operators[rand.Intn(len(operators))]
	firstOperand := rand.Intn(9) + 1

	var secondOperand int
	if op == '/' {
		secondOperand = rand.Intn(9) + 1 // confirm the second operand is not zero
	} else {
		secondOperand = rand.Intn(9)
	}

	expression := fmt.Sprintf("%d %c %d", firstOperand, op, secondOperand)
	fmt.Println(expression)
}
