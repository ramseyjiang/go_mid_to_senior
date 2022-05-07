package main

import "fmt"

var EBread interface{}
var a int
var EVALUE = 666

func main() {
	EBread = EVALUE
	a = EBread.(int)
	fmt.Print(a)
}
