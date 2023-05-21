package bankaccounts

import (
	"errors"
)

type BankAccount struct {
	balance float64
}

func NewBankAccount(initialBalance float64) *BankAccount {
	return &BankAccount{balance: initialBalance}
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount > b.balance {
		return errors.New("insufficient balance for withdrawal")
	}
	b.balance -= amount
	return nil
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

func (b *BankAccount) TransferTo(amount float64, toAccount *BankAccount) error {
	if amount > b.balance {
		return errors.New("insufficient balance for transfer")
	}
	toAccount.Deposit(amount)

	return b.Withdraw(amount)
}
