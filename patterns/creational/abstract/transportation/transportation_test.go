package transportation

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.Build(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.NumWheels())

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorbikeType())

	cruiseMotorbikeVehicle, err := motorbikeF.Build(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("CruiseMotorbikeVehicle vehicle has %d wheels\n", cruiseMotorbikeVehicle.NumWheels())

	cruiseBike, ok := cruiseMotorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Cruise motorbike has type %d\n", cruiseBike.GetMotorbikeType())
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
	t.Logf("Car vehicle has %d seats\n", luxuryCarVehicle.NumWheels())

	luxuryCar, ok := luxuryCarVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())

	familyCarVehicle, err := carF.Build(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats\n", familyCarVehicle.NumWheels())

	familyCar, ok := familyCarVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Family car has %d doors.\n", familyCar.NumDoors())
}
