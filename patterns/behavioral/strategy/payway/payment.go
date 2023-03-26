package payway

import "strconv"

// Strategy Define the interface for the strategy
type Strategy interface {
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
	PaymentStrategy Strategy
}

func (s *ShoppingCart) SetPaymentStrategy(paymentStrategy Strategy) {
	s.PaymentStrategy = paymentStrategy
}

func (s *ShoppingCart) Pay() string {
	return s.PaymentStrategy.Pay(s.Amount)
}
