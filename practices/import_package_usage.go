package practices

import (
	"fmt"

	"golang_learn/customizepkgs/greet"
)

func TriggerGreet() {
	fmt.Println(greet.Morning)
	fmt.Println(greet.SayHi())
}
