package practices

import (
	"fmt"
	"log"
	"reflect"

	"github.com/ramseyjiang/go_mid_to_senior/customizepkgs/keyboard"
)

func TriggerKeyboardInput() {
	// Print doesn’t skip to a new terminal line after printing a message, which lets us keep the prompt and the user’s entry on the same line
	fmt.Print("Enter a grade: ")

	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	// If not declare the status outside the condition, it will have a variable scope error.
	// In go, the variable scope is different with PHP.
	// It has condition scope, function scope, package scope and file scope. So if a variable wants to use outside a condition scope, it must be declared outside the condition scope first.
	// condition scope is also called condition block. Then function scope equals to function block. Others are the same.
	var status string
	if grade >= 60 {
		status = "You passed"
	} else {
		status = "You failed"
	}

	fmt.Println(reflect.TypeOf(grade), grade)
	fmt.Println(status)
}
