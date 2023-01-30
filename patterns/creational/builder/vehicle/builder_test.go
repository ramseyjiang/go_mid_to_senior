package vehicle

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCarBuilder(t *testing.T) {
	manufacturingComplex := ManufactureDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	} else {
		assert.Equal(t, car.Wheels, 4)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	} else {
		assert.Equal(t, car.Structure, "Car")
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	} else {
		assert.Equal(t, car.Seats, 5)
	}
}

func TestBikeBuilder(t *testing.T) {
	manufacturingComplex := ManufactureDirector{}
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	bike := bikeBuilder.GetVehicle()
	// different vehicle has different seats, the default seats are 4, it is for a car.
	// Here is an example to reset seats. You can use the same way to reset wheels also.
	bike.Seats = 1

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 and they were %d\n", bike.Wheels)
	} else {
		assert.Equal(t, bike.Wheels, 2)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Structure on a bike must be 'Bike' and was %s\n", bike.Structure)
	} else {
		assert.Equal(t, bike.Structure, "Bike")
	}
	if bike.Seats != 1 {
		t.Errorf("Seats on a bike must be 2 and they were %d\n", bike.Seats)
	} else {
		assert.Equal(t, bike.Seats, 1)
	}
}

func TestShuttleBusBuilder(t *testing.T) {
	manufacturingComplex := ManufactureDirector{}
	shuttleBusBuilder := &ShuttleBusBuilder{}
	manufacturingComplex.SetBuilder(shuttleBusBuilder)
	manufacturingComplex.Construct()
	shuttleBus := shuttleBusBuilder.GetVehicle()
	// different vehicle has different seats, the default seats are 4, it is for a car.
	// Here is an example to reset seats. You can use the same way to reset wheels also.
	shuttleBus.Seats = 30

	if shuttleBus.Wheels != 8 {
		t.Errorf("Wheels on a shuttleBus must be 8 and they were %d\n", shuttleBus.Wheels)
	} else {
		assert.Equal(t, shuttleBus.Wheels, 8)
	}

	if shuttleBus.Structure != "ShuttleBus" {
		t.Errorf("Structure on a ShuttleBus must be 'ShuttleBus' and was %s\n", shuttleBus.Structure)
	} else {
		assert.Equal(t, shuttleBus.Structure, "ShuttleBus")
	}
	if shuttleBus.Seats != 30 {
		t.Errorf("Seats on a shuttleBus must be 2 and they were %d\n", shuttleBus.Seats)
	} else {
		assert.Equal(t, shuttleBus.Seats, 30)
	}
}
