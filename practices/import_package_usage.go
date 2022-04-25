package practices

import (
	"fmt"

	"github.com/ramseyjiang/go_mid_to_senior/customizepkgs/greet"
)

func TriggerGreet() {
	fmt.Println(greet.Morning)
	fmt.Println(greet.SayHi())
}
