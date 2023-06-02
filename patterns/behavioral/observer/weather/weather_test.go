package weather

import (
	"testing"
)

func TestObserver(t *testing.T) {
	weatherContent := &Content{}
	currentDisplay := &CurrentConditionsDisplay{weatherContent: weatherContent}
	weatherContent.RegisterObserver(currentDisplay)

	tests := []struct {
		name        string
		temperature float64
		humidity    float64
		pressure    float64
	}{
		{"Test1", 80.0, 65.0, 30.4},
		{"Test2", 82.0, 70.0, 29.2},
		{"Test3", 78.0, 90.0, 29.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weatherContent.SetMeasurements(tt.temperature, tt.humidity, tt.pressure)
			if currentDisplay.temperature != tt.temperature || currentDisplay.humidity != tt.humidity || currentDisplay.pressure != tt.pressure {
				t.Errorf("Update() = %v, want %v", currentDisplay.temperature, tt.temperature)
				t.Errorf("Update() = %v, want %v", currentDisplay.humidity, tt.humidity)
				t.Errorf("Update() = %v, want %v", currentDisplay.pressure, tt.pressure)
			}
		})
	}

	// Remove the observer
	weatherContent.RemoveObserver(currentDisplay)

	// Check if the observer has been removed
	if len(weatherContent.observers) != 0 {
		t.Errorf("RemoveObserver() failed, expected length of observers slice to be 0, got %v", len(weatherContent.observers))
	}
}
