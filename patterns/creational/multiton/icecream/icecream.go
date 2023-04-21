package icecream

import (
	"fmt"
	"sync"
)

type Multiton struct {
	Fruit  string
	Favour string
}

var (
	fruitMap = make(map[string]*Multiton)
	mutex    = &sync.Mutex{}
)

func GetIceCream(fruit string) *Multiton {
	mutex.Lock()
	defer mutex.Unlock()

	if fruitMap[fruit] == nil {
		fruitMap[fruit] = &Multiton{
			Fruit:  fruit,
			Favour: fmt.Sprintf("Favour is %s", fruit),
		}
	}

	return fruitMap[fruit]
}
