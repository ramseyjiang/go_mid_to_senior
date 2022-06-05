package keyboardinput

import (
	"log"

	"github.com/ramseyjiang/go_mid_to_senior/customizepkgs/keyboard"
)

func TriggerToCelsius() {
	log.Println("Enter a temperature in Fahrenheit:")

	// Call greeting() to get a temperature.
	fahrenheit, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	log.Printf("%0.2f degrees Celsius\n", celsius)
}
