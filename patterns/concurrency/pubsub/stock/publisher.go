package stock

import (
	"sync"
)

type Publisher interface {
	start()
	AddSubscriber()
	RemoveSubscriber()
	Notify(ticker string, price float64)
	UpdatePrice(ticker string, price float64)
}

// StockExchange Publisher (Stock Exchange)
type StockExchange struct {
	tickers     map[string]float64
	subscribers map[Subscriber]struct{}
	mu          sync.Mutex
}

func NewStockExchange() *StockExchange {
	return &StockExchange{
		tickers:     make(map[string]float64),
		subscribers: make(map[Subscriber]struct{}),
	}
}

func (se *StockExchange) UpdatePrice(ticker string, price float64) {
	se.mu.Lock()
	se.tickers[ticker] = price
	se.mu.Unlock()
	se.Notify(ticker, price)
}

func (se *StockExchange) Notify(ticker string, price float64) {
	se.mu.Lock()
	defer se.mu.Unlock()
	for listener := range se.subscribers {
		listener.OnPriceUpdate(ticker, price)
	}
}

func (se *StockExchange) AddSubscriber(subscriber Subscriber) {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.subscribers[subscriber] = struct{}{}
}
