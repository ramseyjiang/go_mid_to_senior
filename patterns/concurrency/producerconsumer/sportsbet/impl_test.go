package sportsbet

import (
	"testing"
	"time"
)

func TestBetting(t *testing.T) {
	tests := []struct {
		name          string
		order         *BetOrder
		expectedError string
		expectedWin   float64
	}{
		{
			name: "Place valid order",
			order: &BetOrder{
				UserID:  "user123",
				MatchID: "match456",
				BetType: "WIN",
				Amount:  100,
				Odds:    1.5,
			},
			expectedError: "",
			expectedWin:   150,
		},
		{
			name: "Place invalid order",
			order: &BetOrder{
				UserID:  "user123",
				MatchID: "match456",
				BetType: "WIN",
				Amount:  -100,
				Odds:    1.5,
			},
			expectedError: "invalid bet amount",
			expectedWin:   0,
		},
	}

	// Initialize and Start the consumer in a goroutine
	go ProcessOrders()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Handle Synchronization
			err := ReceiveOrders(tt.order)
			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("expected error %v but got %v", tt.expectedError, err.Error())
			} else if err == nil && tt.expectedError != "" {
				t.Errorf("expected error %v but got nil", tt.expectedError)
			}

			if tt.order.PotentialWin != tt.expectedWin {
				t.Errorf("expected potential win to be %v but got %v", tt.expectedWin, tt.order.PotentialWin)
			}
		})
	}

	// Give some time for the consumer to process the orders before ending the tests
	time.Sleep(1 * time.Second)
}
