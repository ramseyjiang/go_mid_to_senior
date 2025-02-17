package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	operators := []string{"+", "-", "*", "/"}
	expression := fmt.Sprintf("%d", rand.Intn(9)+1)
	op := operators[rand.Intn(len(operators))]
	num := rand.Intn(9) + 1
	expression += fmt.Sprintf(" %s %d", op, num)
	fmt.Println(expression)
}
