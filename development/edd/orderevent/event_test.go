package orderevent

import (
	"fmt"
	"testing"
)

func TestProcessEvent(t *testing.T) {
	order := Order{
		ID: "1234",
		Items: map[string]int{
			"item1": 1,
			"item2": 2,
		},
	}

	tests := []struct {
		name     string
		event    Event
		expected string
	}{
		{
			name:     "OrderPlacedEvent",
			event:    OrderPlacedEvent{Order: order},
			expected: fmt.Sprintf("Order %s has been placed\n", order.ID),
		},
		{
			name:     "InventoryReservedEvent",
			event:    InventoryReservedEvent{Order: order},
			expected: fmt.Sprintf("Inventory has been reserved for order %s\n", order.ID),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.event.ProcessEvent()
			if result != test.expected {
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})
	}
}
