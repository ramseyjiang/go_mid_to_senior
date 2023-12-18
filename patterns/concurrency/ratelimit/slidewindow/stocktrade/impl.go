package stocktrade

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type TradeWindow struct {
	Trades []time.Time
	Mutex  sync.Mutex
}

func NewTradeWindow() *TradeWindow {
	return &TradeWindow{
		Trades: []time.Time{},
	}
}

func (tw *TradeWindow) AddTrade(t time.Time) {
	tw.Mutex.Lock()
	defer tw.Mutex.Unlock()
	tw.Trades = append(tw.Trades, t)
}

func (tw *TradeWindow) AllowTrade() bool {
	tw.Mutex.Lock()
	defer tw.Mutex.Unlock()

	now := time.Now()
	oneMinuteAgo := now.Add(-1 * time.Minute)

	// Filter trades to keep only those within the last minute
	validTrades := []time.Time{}
	for _, tradeTime := range tw.Trades {
		if tradeTime.After(oneMinuteAgo) {
			validTrades = append(validTrades, tradeTime)
		}
	}
	tw.Trades = validTrades

	// Allow trade if there are less than 5 trades in the last minute
	return len(tw.Trades) < 5
}

func TradeLimitMiddleware(tw *TradeWindow, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !tw.AllowTrade() {
			http.Error(w, "Trade limit exceeded", http.StatusTooManyRequests)
			return
		}

		tw.AddTrade(time.Now())
		next.ServeHTTP(w, r)
	})
}

func TradeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Trade executed successfully")
	// Do exact logic after trade executed.
}
