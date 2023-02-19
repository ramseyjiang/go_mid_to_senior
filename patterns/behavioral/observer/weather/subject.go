package weather

// Subject interface
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// Station is a Concrete Subject
type Station struct {
	observers   []Observer
	temperature float64
}

func (w *Station) RegisterObserver(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *Station) RemoveObserver(observer Observer) {
	for i, o := range w.observers {
		if o == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *Station) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update(w)
	}
}

func (w *Station) SetTemperature(temperature float64) {
	w.temperature = temperature
	w.NotifyObservers()
}
