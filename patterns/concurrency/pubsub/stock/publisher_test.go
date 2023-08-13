package stock

import (
	"testing"
)

func TestStockExchange(t *testing.T) {
	nasdaq := NewStockExchange()
	nyse := NewStockExchange()

	// Testing multiple price updates and notifications across exchanges
	t.Run("MultipleUpdatesAndNotify", func(t *testing.T) {
		mockSub1 := &MockSubscriber{}
		mockSub2 := &MockSubscriber{}

		nasdaq.AddSubscriber(mockSub1)
		nasdaq.AddSubscriber(mockSub2)
		nyse.AddSubscriber(mockSub2)

		nasdaq.UpdatePrice("AAPL", 150.0)
		nyse.UpdatePrice("GOOGL", 2500.0)

		if mockSub1.LastPrice != 150.0 || mockSub1.LastTicker != "AAPL" {
			t.Fatalf("MockSub1 expected AAPL with 150.0, got %s with %.2f", mockSub1.LastTicker, mockSub1.LastPrice)
		}

		if mockSub2.LastPrice != 2500.0 || mockSub2.LastTicker != "GOOGL" {
			t.Fatalf("MockSub2 expected GOOGL with 2500.0, got %s with %.2f", mockSub2.LastTicker, mockSub2.LastPrice)
		}
	})
}

// MockSubscriber for testing purpose, similar to the one you have seen earlier.
type MockSubscriber struct {
	LastTicker string
	LastPrice  float64
}

func (ms *MockSubscriber) OnPriceUpdate(ticker string, price float64) {
	ms.LastTicker = ticker
	ms.LastPrice = price
}
