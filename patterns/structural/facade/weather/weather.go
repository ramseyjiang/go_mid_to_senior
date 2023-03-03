package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	commonRequestPrefix              = "http://api.openweathermap.org/data/2.5/"
	weatherByCityName                = commonRequestPrefix + "weather?q=%s,%s&APPID=%s"
	weatherByGeographicalCoordinates = commonRequestPrefix + "weather?lat=%f&lon=%f&APPID=%s"
)

// DataRetriever is a Facade interface
type DataRetriever interface {
	GetByGeoCoordinates(lat, lon float32) (*Weather, error)
	GetByCityAndCountryCode(city, countryCode string) (*Weather, error)
}

// CurrentWeatherData is used to implement the Facade interface using a concrete implementation
type CurrentWeatherData struct {
	AppID string `json:"-"`
}

// Weather is used to define the response structure.
type Weather struct {
	Coordinate struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coordinate"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

// GetGeoCoordinates returns the current weather data by passing a geographical coordinates (latitude and longitude) in decimal notation.
// It returns weather information or a detailed error.
// For example, to query about Madrid, Spain you get: currentWeather.GetGeoCoordinates(-3, 40)
func (c *CurrentWeatherData) GetGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf(weatherByGeographicalCoordinates, lat, lon, c.AppID))
}

// GetCityCountryCode returns the current weather data by passing a city name and an ISO country code.
// It returns weather information or a detailed error
// For example, to query about Madrid, Spain you get: currentWeather.GetCityCountryCode("Madrid", "ES)
func (c *CurrentWeatherData) GetCityCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf(weatherByCityName, city, countryCode, c.AppID))
}

func (c *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}

	return w, nil
}

// doRequest is used to define the APIs, it belongs to the step implement the complex subsystem's classes, interfaces, and APIs.
func (c *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		byt, errMsg := io.ReadAll(resp.Body)
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was:\n%s\n",
			resp.StatusCode, errMsg)

		return
	}

	weather, err = c.responseParser(resp.Body)
	if err != nil {
		return
	}
	_ = resp.Body.Close()

	return
}
