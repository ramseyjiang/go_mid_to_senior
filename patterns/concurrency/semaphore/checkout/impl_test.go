package checkout

import (
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name    string
		orderID int
	}{
		{"Order 1", 201},
		{"Order 2", 202},
		{"Order 3", 203},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg.Add(1)
			go Process(tt.orderID)
		})
	}

	wg.Wait()
}

func TestProcessOrders(t *testing.T) {
	orderIDs := []int{301, 302, 303, 304, 305}
	ProcessOrders(orderIDs)
}
