package icecream

import (
	"fmt"
	"sync"
)

type Multiton struct {
	Fruit  string
	Favour string
}

var fruitMap map[string]*Multiton
var once sync.Once

func GetIceCream(fruit string) *Multiton {
	once.Do(func() {
		fruitMap = make(map[string]*Multiton)
	})

	if fruitMap[fruit] == nil {
		fruitMap[fruit] = &Multiton{
			Fruit:  fruit,
			Favour: fmt.Sprintf("Favour is %s", fruit),
		}
	}

	return fruitMap[fruit]
}
