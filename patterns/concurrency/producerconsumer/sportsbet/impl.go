package sportsbet

import "github.com/pkg/errors"

type BetOrder struct {
	UserID       string
	MatchID      string
	BetType      string // e.g. "WIN", "LOSE", "DRAW"
	Amount       float64
	Odds         float64
	PotentialWin float64
}

var orderChannel = make(chan *BetOrder, 10) // Channel to hold up to 10 orders

// ReceiveOrders is a Producer Function to receive orders
func ReceiveOrders(order *BetOrder) error {
	if order.Amount <= 0 {
		return errors.New("invalid bet amount")
	}
	order.PotentialWin = order.Amount * order.Odds
	orderChannel <- order
	return nil
}

// ProcessOrders func is a Consumer Function to process orders:
func ProcessOrders() {
	for order := range orderChannel {
		// Here, you can add logic to save the order to a database or any other storage.
		// For this example, we'll just print the order.
		println("Processed order for user:", order.UserID)
	}
}
