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
