package cs

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGame(t *testing.T) {
	gameTest := newGame()

	// Add Terrorist
	gameTest.addTerrorist(TerroristDressType)
	gameTest.addTerrorist(TerroristDressType)
	gameTest.addTerrorist(TerroristDressType)
	gameTest.addTerrorist(TerroristDressType)

	// Add CounterTerrorist
	gameTest.addCounterTerrorist(CounterTerroristDressType)
	gameTest.addCounterTerrorist(CounterTerroristDressType)
	gameTest.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		if dressType == "tDress" {
			assert.Equal(t, "red", dress.getColor())
		}

		if dressType == "ctDress" {
			assert.Equal(t, "green", dress.getColor())
		}
	}
}
