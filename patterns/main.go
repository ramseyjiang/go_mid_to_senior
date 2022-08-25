package main

import (
	"fmt"

	"github.com/ramseyjiang/go_mid_to_senior/patterns/creationptns/factory"
	"github.com/ramseyjiang/go_mid_to_senior/patterns/creationptns/singleton"
)

func main() {
	fmt.Println("print factory pattern output start -------")
	factory.Entry()
	fmt.Println("print factory pattern output end ---------")

	fmt.Println("\nprint singleton pattern output start -----")
	singleton.GetInstance().SetTitle("test set title")
	fmt.Println(singleton.GetInstance().GetTitle())
	fmt.Println("print singleton pattern output end -------")
}
