package payment

import "fmt"

type DebitCardPM struct{}

func (dc *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card", amount)
}
