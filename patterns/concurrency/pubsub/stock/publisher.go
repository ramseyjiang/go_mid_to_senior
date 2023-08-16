package stock

import (
	"sync"
)

type Publisher interface {
	AddSubscriber(s Subscriber)
	RemoveSubscriber(s Subscriber)
	Notify(ticker string, price float64)
	UpdatePrice(ticker string, price float64)
}

type StockPublisher struct {
	tickers     map[string]float64
	subscribers map[Subscriber]struct{}
	mu          sync.Mutex
}

func NewStockExchange() *StockPublisher {
	return &StockPublisher{
		tickers:     make(map[string]float64),
		subscribers: make(map[Subscriber]struct{}),
	}
}

func (sp *StockPublisher) UpdatePrice(ticker string, price float64) {
	sp.mu.Lock()
	sp.tickers[ticker] = price
	sp.mu.Unlock()
	sp.Notify(ticker, price)
}

func (sp *StockPublisher) Notify(ticker string, price float64) {
	sp.mu.Lock()
	defer sp.mu.Unlock()
	for listener := range sp.subscribers {
		listener.OnPriceUpdate(ticker, price)
	}
}

// AddSubscriber adds a subscriber to the StockPublisher.
func (sp *StockPublisher) AddSubscriber(s Subscriber) {
	sp.mu.Lock()
	defer sp.mu.Unlock()
	sp.subscribers[s] = struct{}{}
}

// RemoveSubscriber removes a subscriber from the StockPublisher.
func (sp *StockPublisher) RemoveSubscriber(s Subscriber) {
	sp.mu.Lock()
	defer sp.mu.Unlock()
	delete(sp.subscribers, s)
}
