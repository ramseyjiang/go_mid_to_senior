package weather

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestObserver(t *testing.T) {
	weatherStation := &Station{}
	phoneDisplay := &PhoneDisplay{}
	pcDisplay := &PCDisplay{}

	weatherStation.RegisterObserver(phoneDisplay)
	weatherStation.RegisterObserver(pcDisplay)

	weatherStation.SetTemperature(25.5)
	assert.Equal(t, 25.5, phoneDisplay.GetTemp())
	assert.Equal(t, 25.5, pcDisplay.GetTemp())

	weatherStation.RemoveObserver(phoneDisplay)
	weatherStation.SetTemperature(30.2)
	assert.Equal(t, 30.2, pcDisplay.GetTemp())
	assert.NotEqual(t, 30.2, phoneDisplay.GetTemp()) // because the phoneDisplay observer was removed

	weatherStation.RemoveObserver(pcDisplay) // it should remove observer first, if not, the observer will also be updated.
	weatherStation.SetTemperature(24.81)
	assert.NotEqual(t, 24.81, phoneDisplay.GetTemp()) // because the phoneDisplay observer was removed
	assert.NotEqual(t, 24.81, pcDisplay.GetTemp())
}
