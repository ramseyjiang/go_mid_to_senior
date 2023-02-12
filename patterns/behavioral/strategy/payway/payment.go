package payway

import "strconv"

// In this example, the interface PaymentStrategy defines the methods for the different payment strategies that can be used by the ShoppingCart class. The PaypalStrategy and CreditCardStrategy structs both implement the PaymentStrategy interface and define the specific behavior for each strategy.
//
// The ShoppingCart struct contains a reference to the PaymentStrategy interface, and the SetPaymentStrategy method allows you to change the payment strategy at runtime. Finally, the Pay method calls the Pay method of the currently set payment strategy to make the payment.

// PaymentStrategy Define the interface for the strategy
type PaymentStrategy interface {
	Pay(amount int) string
}

// PaypalStrategy Concrete implementations of the strategy interface
type PaypalStrategy struct{}

func (p *PaypalStrategy) Pay(amount int) string {
	return "Paying " + strconv.Itoa(amount) + " using PayPal"
}

type CreditCardStrategy struct{}

func (c *CreditCardStrategy) Pay(amount int) string {
	return "Paying " + strconv.Itoa(amount) + " using Credit Card"
}

// ShoppingCart The context class
type ShoppingCart struct {
	Amount          int
	PaymentStrategy PaymentStrategy
}

func (s *ShoppingCart) SetPaymentStrategy(paymentStrategy PaymentStrategy) {
	s.PaymentStrategy = paymentStrategy
}

func (s *ShoppingCart) Pay() string {
	return s.PaymentStrategy.Pay(s.Amount)
}
