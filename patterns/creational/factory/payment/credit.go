package payment

import "fmt"

type CreditCardPM struct{}

func (cc *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using credit card", amount)
}
