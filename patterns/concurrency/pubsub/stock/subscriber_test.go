package stock

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestTradingBot(t *testing.T) {
	bot := &TradingBot{ID: "testBot"}

	t.Run("MultiplePublisherUpdates", func(t *testing.T) {
		nasdaq := NewStockExchange()
		nyse := NewStockExchange()

		nasdaq.AddSubscriber(bot)
		nyse.AddSubscriber(bot)

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

		if !strings.Contains(output, "TradingBot testBot: Received AAPL price update: $150.00") {
			t.Errorf("Expected output for AAPL update not found")
		}

		if !strings.Contains(output, "TradingBot testBot: Received GOOGL price update: $2500.00") {
			t.Errorf("Expected output for GOOGL update not found")
		}
	})
}
