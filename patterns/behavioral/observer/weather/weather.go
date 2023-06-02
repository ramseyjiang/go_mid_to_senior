package weather

import "fmt"

// Step 1: Define a subject interface and an observer interface.

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type Observer interface {
	Update(temp float64, humidity float64, pressure float64)
}

// Step 2: Implement the subject interface.

type Content struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func (c *Content) RegisterObserver(o Observer) {
	c.observers = append(c.observers, o)
}

func (c *Content) RemoveObserver(o Observer) {
	for i, observer := range c.observers {
		if observer == o {
			// Remove the observer from the slice.
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

func (c *Content) NotifyObservers() {
	for _, observer := range c.observers {
		observer.Update(c.temperature, c.humidity, c.pressure)
	}
}

func (c *Content) MeasurementsChanged() {
	c.NotifyObservers()
}

func (c *Content) SetMeasurements(temperature float64, humidity float64, pressure float64) {
	c.temperature = temperature
	c.humidity = humidity
	c.pressure = pressure
	c.MeasurementsChanged()
}

// Step 3:Implement the observer interface.

type CurrentConditionsDisplay struct {
	temperature    float64
	humidity       float64
	pressure       float64
	weatherContent Subject
}

func (ccd *CurrentConditionsDisplay) Update(temp float64, humidity float64, pressure float64) {
	ccd.temperature = temp
	ccd.humidity = humidity
	ccd.pressure = pressure
	ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.2f F degrees and %.2f humidity\n", ccd.temperature, ccd.humidity)
}
