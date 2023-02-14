package payway

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPayment(t *testing.T) {
	// Use the context class and the strategies
	shoppingCart := &ShoppingCart{Amount: 100}

	// Set the payment strategy to PayPal
	shoppingCart.SetPaymentStrategy(&PaypalStrategy{})
	assert.Equal(t, "Paying 100 using PayPal", shoppingCart.Pay())

	// Set the payment strategy to Credit Card
	shoppingCart.SetPaymentStrategy(&CreditCardStrategy{})
	assert.Equal(t, "Paying 100 using Credit Card", shoppingCart.Pay())
}
