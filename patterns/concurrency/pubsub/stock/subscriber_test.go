package stock

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestTradingBots(t *testing.T) {
	bot1 := &TradingBot{ID: "bot1"}
	bot2 := &TradingBot{ID: "bot2"}

	nasdaq := NewStockExchange()
	nyse := NewStockExchange()

	t.Run("MultiplePublishersAndSubscribers", func(t *testing.T) {
		// Add both bots as subscribers to both exchanges
		nasdaq.AddSubscriber(bot1)
		nasdaq.AddSubscriber(bot2)
		nyse.AddSubscriber(bot1)
		nyse.AddSubscriber(bot2)

		// Capturing stdout to a buffer
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		nasdaq.UpdatePrice("AAPL", 150.0)
		nyse.UpdatePrice("GOOGL", 2500.0)

		w.Close()
		os.Stdout = old

		var buf bytes.Buffer
		buf.ReadFrom(r)
		output := buf.String()

		expectedOutputs := []string{
			"TradingBot bot1: Received AAPL price update: $150.00",
			"TradingBot bot2: Received AAPL price update: $150.00",
			"TradingBot bot1: Received GOOGL price update: $2500.00",
			"TradingBot bot2: Received GOOGL price update: $2500.00",
		}

		for _, expected := range expectedOutputs {
			if !strings.Contains(output, expected) {
				t.Errorf("Expected output \"%s\" not found", expected)
			}
		}
	})
}
