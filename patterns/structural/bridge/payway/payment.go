package payway

import "fmt"

// PaymentMethod is Implementor interface
type PaymentMethod interface {
	Pay(amount float64) string
}

// CreditCard is Concrete Implementor A
type CreditCard struct {
	Name string
}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("%s paid %.2f using credit card", c.Name, amount)
}

// PayPal is Concrete Implementor B
type PayPal struct {
	Name string
}

func (p *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("%s paid %.2f using PayPal", p.Name, amount)
}
