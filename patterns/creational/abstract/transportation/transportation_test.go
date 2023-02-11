package transportation

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	sportMotorbike, err := motorbikeF.Build(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, sportMotorbike.NumWheels())

	sportBike, ok := sportMotorbike.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	assert.Equal(t, 1, sportBike.GetMotorbikeType())

	cruiseMotorbike, err := motorbikeF.Build(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, cruiseMotorbike.NumWheels())

	cruiseBike, ok := cruiseMotorbike.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	assert.Equal(t, 2, cruiseBike.GetMotorbikeType())
}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	luxuryCarVehicle, err := carF.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 5, luxuryCarVehicle.NumSeats())

	luxuryCar, ok := luxuryCarVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	assert.Equal(t, 4, luxuryCar.NumDoors())

	familyCarVehicle, err := carF.Build(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 5, familyCarVehicle.NumSeats())

	familyCar, ok := familyCarVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	assert.Equal(t, 5, familyCar.NumDoors())
}
