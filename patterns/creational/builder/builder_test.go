package builder

import "testing"

func TestCarBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}
}

func TestBikeBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	bike := bikeBuilder.GetVehicle()
	bike.Seats = 1
	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 and they were %d\n", bike.Wheels)
	}
	if bike.Structure != "Bike" {
		t.Errorf("Structure on a bike must be 'Bike' and was %s\n", bike.Structure)
	}
	if bike.Seats != 1 {
		t.Errorf("Seats on a bike must be 2 and they were %d\n", bike.Seats)
	}
}
