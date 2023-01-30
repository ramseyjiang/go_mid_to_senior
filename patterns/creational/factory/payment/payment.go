package payment

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	Cash       = 1
	DebitCard  = 2
	CreditCard = 3
)

type PayWay interface {
	Pay(amount float32) string
}

func GetPayWay(m int) (PayWay, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized", m))
	}
}

type CashPM struct{}
type DebitCardPM struct{}
type CreditCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash", amount)
}
func (dc *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card", amount)
}

func (cc *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using credit card", amount)
}
