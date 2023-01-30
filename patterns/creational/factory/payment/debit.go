package payment

import "fmt"

type DebitPay struct {
	PayInfo
}

func newDebit() PayWay {
	return &DebitPay{
		PayInfo: PayInfo{
			name: DebitName,
		},
	}
}

func (d *DebitPay) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card", amount)
}
