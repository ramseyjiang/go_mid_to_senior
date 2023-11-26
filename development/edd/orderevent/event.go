package orderevent

import (
	"fmt"
)

// Order data structure
type Order struct {
	ID    string
	Items map[string]int // map of item ID and quantity
}

type Event interface {
	ProcessEvent() string
}

type OrderPlacedEvent struct {
	Order Order
}

func (e OrderPlacedEvent) ProcessEvent() string {
	message := fmt.Sprintf("Order %s has been placed\n", e.Order.ID)
	// Here you can do things like reserving inventory, notifying the user, etc.
	return message
}

type InventoryReservedEvent struct {
	Order Order
}

func (e InventoryReservedEvent) ProcessEvent() string {
	message := fmt.Sprintf("Inventory has been reserved for order %s\n", e.Order.ID)
	// Here you can do things like preparing the order for shipping, etc.
	return message
}
