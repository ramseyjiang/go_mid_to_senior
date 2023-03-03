package weather

import (
	"bytes"
	"io"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetCityCountryCode(t *testing.T) {
	r := getMockData()
	weatherMap := CurrentWeatherData{AppID: ""}

	// Implement the client that uses the Facade to interact with the subsystem.
	weather, _ := weatherMap.GetCityCountryCode("Madrid", "ES")
	mockWeather, err := weatherMap.responseParser(r) // using mock data replace the weather return
	if err != nil {
		t.Fatal(err)
	}
	weather = mockWeather

	assert.Equal(t, float32(-3.7), weather.Coordinate.Lon)
	assert.Equal(t, float32(40.42), weather.Coordinate.Lat)
}

func TestGetGeoCoordinates(t *testing.T) {
	r := getMockData()
	weatherMap := CurrentWeatherData{AppID: ""}

	// Implement the client that uses the Facade to interact with the subsystem.
	weather, _ := weatherMap.GetGeoCoordinates(-3.7, 40.42)
	mockWeather, err := weatherMap.responseParser(r) // // using mock data replace the weather return
	if err != nil {
		t.Fatal(err)
	}
	weather = mockWeather

	assert.Equal(t, 200, weather.Cod)
	if weather.Cod != 200 {
		t.Errorf("Cod was not 200 as expected. Code: %d\n", weather.Cod)
	}
}

func getMockData() io.Reader {
	response := `{"coordinate":{"lon":-3.7,"lat":40.42},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],"base":"stations","id":3117735,"name":"Madrid","cod":200}`
	r := bytes.NewReader([]byte(response))

	return r
}

func TestResponseParser(t *testing.T) {
	r := getMockData()
	weatherMap := CurrentWeatherData{AppID: ""}

	weather, err := weatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3117735, weather.ID)
}
