package payment

import "fmt"

type CashPay struct {
	PayInfo
}

func newCash() PayWay {
	return &CashPay{
		PayInfo: PayInfo{
			name: CashName,
		},
	}
}

func (c *CashPay) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash", amount)
}
