package main

import "fmt"

var EBread interface{}
var a int
var EValue = 666

func main() {
	EBread = EValue
	a = EBread.(int)
	fmt.Print(a)
}
