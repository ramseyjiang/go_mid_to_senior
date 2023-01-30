package payment

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	Cash       = 1
	DebitCard  = 2
	CreditCard = 3
	CashName   = "cash"
	DebitName  = "debit"
	CreditName = "credit"
)

type PayWay interface {
	Pay(amount float32) string
	SetName(name string)
	GetName() string
}

type PayInfo struct {
	name string
}

func (p *PayInfo) SetName(name string) {
	p.name = name
}

func (p *PayInfo) GetName() string {
	return p.name
}

func GetPayWay(payType int) (PayWay, error) {
	switch payType {
	case Cash:
		return newCash(), nil
	case DebitCard:
		return newDebit(), nil
	case CreditCard:
		return newCredit(), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized", payType))
	}
}
