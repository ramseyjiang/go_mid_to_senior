package payway

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPayment(t *testing.T) {
	item1 := &Item{
		Name:    "Item 1",
		Price:   19.99,
		Payment: &CreditCard{"John Doe"},
	}
	assert.Equal(t, "John Doe paid 19.99 using credit card", item1.Purchase())

	item2 := &Item{
		Name:    "Item 2",
		Price:   99.99,
		Payment: &PayPal{"Jane Doe"},
	}
	assert.Equal(t, "Jane Doe paid 99.99 using PayPal", item2.Purchase())
}
