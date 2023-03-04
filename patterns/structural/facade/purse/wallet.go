package purse

import (
	"github.com/pkg/errors"
)

// WalletOperation is the complex subsystem's interfaces
type WalletOperation interface {
	creditBalance(amount int)
	debitBalance(amount int) error
}

type Wallet struct {
	balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return errors.New("balance is not sufficient")
	}

	w.balance = w.balance - amount
	return nil
}
