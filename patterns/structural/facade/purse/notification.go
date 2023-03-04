package purse

// Notice is the complex subsystem's interfaces
type Notice interface {
	SendWalletCreditNotice() string
	SendWalletDebitNotice() string
}

type Notification struct {
	msg string
}

func (n *Notification) SendWalletCreditNotice() {
	n.msg = "Sending wallet credit notification"
}

func (n *Notification) SendWalletDebitNotice() {
	n.msg = "Sending wallet debit notification"
}
