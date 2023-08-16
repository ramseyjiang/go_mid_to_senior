package stock

import (
	"testing"
)

func TestStockExchange(t *testing.T) {
	// Create two publishers
	nasdaq := NewStockExchange()
	nyse := NewStockExchange()

	// Create two subscribers (TradingBots)
	bot1 := &TradingBot{ID: "Bot1"}
	bot2 := &TradingBot{ID: "Bot2"}

	// Test: Adding subscribers to publishers
	t.Run("AddSubscribers", func(t *testing.T) {
		nasdaq.AddSubscriber(bot1)
		nasdaq.AddSubscriber(bot2)
		nyse.AddSubscriber(bot1)
		nyse.AddSubscriber(bot2)

		if _, exists := nasdaq.subscribers[bot1]; !exists {
			t.Errorf("Bot1 not added to nasdaq")
		}
		if _, exists := nasdaq.subscribers[bot2]; !exists {
			t.Errorf("Bot2 not added to nasdaq")
		}
		if _, exists := nyse.subscribers[bot1]; !exists {
			t.Errorf("Bot1 not added to nyse")
		}
		if _, exists := nyse.subscribers[bot2]; !exists {
			t.Errorf("Bot2 not added to nyse")
		}
	})

	// Test: Removing subscribers from publishers
	t.Run("RemoveSubscribers", func(t *testing.T) {
		nasdaq.RemoveSubscriber(bot1)
		nyse.RemoveSubscriber(bot2)

		if _, exists := nasdaq.subscribers[bot1]; exists {
			t.Errorf("Bot1 not removed from nasdaq")
		}
		if _, exists := nyse.subscribers[bot2]; exists {
			t.Errorf("Bot2 not removed from nyse")
		}
	})
}
