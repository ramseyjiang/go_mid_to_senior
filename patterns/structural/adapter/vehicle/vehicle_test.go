package vehicle

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestVehicle(t *testing.T) {
	car := &Car{Name: "Car"}
	plane := &Plane{Name: "Plane"}
	boat := &Boat{Name: "Boat"}

	carAdapter := &CarAdapter{car: car}
	planeAdapter := &PlaneAdapter{plane: plane}
	boatAdapter := &BoatAdapter{boat: boat}

	assert.Equal(t, "Car is driving on road.", carAdapter.Drive())
	assert.Equal(t, "Plane is flying in sky.", planeAdapter.Drive())
	assert.Equal(t, "Boat is sailing on water.", boatAdapter.Drive())
}
