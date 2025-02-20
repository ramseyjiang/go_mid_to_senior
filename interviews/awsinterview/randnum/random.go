package main

import (
	"fmt"
	"math/rand"
)

func main() {
	operators := []rune{'+', '-', '*', '/'}
	op := operators[rand.Intn(len(operators))]
	expression := fmt.Sprintf("%d %c %d", rand.Intn(9), op, rand.Intn(9))
	fmt.Println(expression)
}
