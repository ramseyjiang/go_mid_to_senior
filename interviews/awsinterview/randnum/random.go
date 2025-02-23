package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	expression := generateExpression()
	fmt.Println(expression)
}

func generateExpression() string {
	operators := []rune{'+', '-', '*', '/'}
	op := operators[rand.Intn(len(operators))]
	firstOperand := rand.Intn(9) + 1

	var secondOperand int
	if op == '/' {
		secondOperand = rand.Intn(9) + 1 // confirm the second operand is not zero
	} else {
		secondOperand = rand.Intn(9)
	}
	return fmt.Sprintf("%d %c %d", firstOperand, op, secondOperand)
}
