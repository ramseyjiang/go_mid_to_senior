package eventbus

type EventBus struct {
	subscribers map[string][]func(string)
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(string)),
	}
}

func (eb *EventBus) Subscribe(eventType string, callback func(string)) {
	eb.subscribers[eventType] = append(eb.subscribers[eventType], callback)
}

func (eb *EventBus) Publish(eventType string, data string) {
	for _, callback := range eb.subscribers[eventType] {
		callback(data)
	}
}
