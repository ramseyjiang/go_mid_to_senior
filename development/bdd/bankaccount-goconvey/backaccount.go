package bankaccount

import "github.com/pkg/errors"

type BankAccount struct {
	balance int
}

func NewBankAccount(initialBalance int) *BankAccount {
	return &BankAccount{balance: initialBalance}
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount int) error {
	if amount > b.balance {
		return errors.New("Insufficient balance for withdrawal")
	}
	b.balance -= amount
	return nil
}

func (b *BankAccount) Balance() int {
	return b.balance
}
