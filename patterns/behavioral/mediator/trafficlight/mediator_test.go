package trafficlight

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLights(t *testing.T) {
	stationManager := &LightsManager{
		carPass: true,
	}

	train := &Train{
		mediator: stationManager,
	}

	car := &Cars{
		mediator: stationManager,
	}

	t.Run("Train is coming, cars wait", func(t *testing.T) {
		assert.Equal(t, "Train: is coming, cars wait.", train.coming())
		assert.Equal(t, "Cars: Please wait, train is coming", car.coming())
	})

	t.Run("Train is passed, cars can pass", func(t *testing.T) {
		assert.Equal(t, "Train: Left", train.passed())
		assert.Equal(t, "Cars: No train, cars can pass", car.coming())
	})

	t.Run("Train is passing, cars cannot pass", func(t *testing.T) {
		assert.Equal(t, "Train: No car, train is passing", train.passing())
		assert.Equal(t, "Cars: Please wait, train is coming", car.coming())
	})
}
