package payway

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// create instances of the CreditCardProcessor and PayPalProcessor classes and use them to process online orders using the OnlineOrder class.
func TestPayment(t *testing.T) {
	// Process an online order with credit card payment
	creditCardProcessor := &CreditCardProcessor{}
	onlineOrderWithCreditCard := &OnlineOrder{paymentProcessor: creditCardProcessor}
	cardResult := onlineOrderWithCreditCard.ProcessOrder(50.0)
	assert.Equal(t, true, cardResult)

	// Process an online order with PayPal payment
	payPalProcessor := &PayPalProcessor{}
	onlineOrderWithPayPal := &OnlineOrder{paymentProcessor: payPalProcessor}
	paypalResult := onlineOrderWithPayPal.ProcessOrder(75.0)
	assert.Equal(t, true, paypalResult)
}
