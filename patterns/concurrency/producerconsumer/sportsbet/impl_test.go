package sportsbet

import (
	"sync"
	"testing"
)

func TestSportsBet(t *testing.T) {
	var wg sync.WaitGroup

	// Start the consumer in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		ProcessOrders()
	}()

	t.Run("Place valid order", func(t *testing.T) {
		order := &BetOrder{
			UserID:  "user123",
			MatchID: "match456",
			BetType: "WIN",
			Amount:  100,
			Odds:    1.5,
		}
		err := ReceiveOrders(order)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		if order.PotentialWin != 150 {
			t.Errorf("expected potential win to be 150 but got %v", order.PotentialWin)
		}
	})

	t.Run("Place invalid order", func(t *testing.T) {
		order := &BetOrder{
			UserID:  "user123",
			MatchID: "match456",
			BetType: "WIN",
			Amount:  -100,
			Odds:    1.5,
		}
		err := ReceiveOrders(order)
		if err == nil {
			t.Error("expected error but got none")
		}
	})

	// Close the channel and wait for the consumer to finish processing
	close(orderChannel)
	wg.Wait()
}
