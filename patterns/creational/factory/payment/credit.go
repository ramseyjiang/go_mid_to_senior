package payment

import "fmt"

type CreditPay struct {
	PayInfo
}

func newCredit() PayWay {
	return &CreditPay{
		PayInfo: PayInfo{
			name: CreditName,
		},
	}
}

func (c *CreditPay) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using credit card", amount)
}
