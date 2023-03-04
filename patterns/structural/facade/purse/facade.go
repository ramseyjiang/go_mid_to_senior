package purse

// WalletIFace is used to define the Facade interface that provides a simplified interface to the subsystem.
type WalletIFace interface {
	AddMoneyToWallet(accountID string, securityCode int, amount int) error
	DeductMoneyFromWallet(accountID string, securityCode int, amount int)
}

// WalletFacade is used to implement the Facade interface using a concrete implementation that delegates to the subsystem.
type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	walletFacade := &WalletFacade{
		account:      NewAccount(accountID),
		securityCode: NewSecurityCode(code),
		wallet:       NewWallet(),
		notification: &Notification{},
		ledger:       &Ledger{},
	}

	return walletFacade
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.SendWalletCreditNotice()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.SendWalletDebitNotice()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}
