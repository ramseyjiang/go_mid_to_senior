package house

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUsualBuilder(t *testing.T) {
	usualBuilder := getBuilder(usual)
	director := newDirector(usualBuilder)
	usualHouse := director.buildHouse()

	if usualHouse.floor != 2 {
		t.Errorf("floor in a usual house must be 2 and they were %d\n", usualHouse.floor)
	} else {
		assert.Equal(t, usualHouse.floor, 2)
	}

	if usualHouse.windowType != "Wooden Window" {
		t.Errorf("windowType in a usual house must be Wooden Window and they were %s\n", usualHouse.windowType)
	} else {
		assert.Equal(t, usualHouse.floor, 2)
	}

	if usualHouse.doorType != "Wooden Door" {
		t.Errorf("floor in a usual house must be Wooden Door and they were %s\n", usualHouse.doorType)
	} else {
		assert.Equal(t, usualHouse.floor, 2)
	}
}

func TestIglooBuilder(t *testing.T) {
	iglooBuilder := getBuilder(igloo)
	director := newDirector(iglooBuilder)
	iglooHouse := director.buildHouse()

	iglooHouse.floor = 3
	if iglooHouse.floor != 3 {
		t.Errorf("floor in a igloo house must be 2 and they were %d\n", iglooHouse.floor)
	} else {
		assert.Equal(t, iglooHouse.floor, 3)
	}

	if iglooHouse.windowType != "Snow Window" {
		t.Errorf("windowType in a igloo house must be Snow Window and they were %s\n", iglooHouse.windowType)
	} else {
		assert.Equal(t, iglooHouse.windowType, "Snow Window")
	}

	if iglooHouse.doorType != "Snow Door" {
		t.Errorf("floor in a igloo house must be Snow Door and they were %s\n", iglooHouse.doorType)
	} else {
		assert.Equal(t, iglooHouse.doorType, "Snow Door")
	}
}
