package integertoroman

import (
	"log"
	"testing"
)

func TestIntegerToRoman(t *testing.T) {
	num := 2023
	target := "MMXXIII"
	if res := IntToRoman(num); res == target {
		log.Println("num equals to roman ", target)
	} else {
		t.Error("Target is not any two num sum in the nums.")
	}
}
