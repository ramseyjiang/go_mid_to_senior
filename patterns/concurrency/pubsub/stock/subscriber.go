package stock

import "fmt"

// Subscriber (Trading Algorithm/Entity)
type Subscriber interface {
	OnPriceUpdate(string, float64)
}

type TradingBot struct {
	ID string
}

func (tb *TradingBot) OnPriceUpdate(ticker string, price float64) {
	fmt.Printf("TradingBot %s: Received %s price update: $%.2f\n", tb.ID, ticker, price)
	// Trading logic can be implemented here based on price updates.
}
