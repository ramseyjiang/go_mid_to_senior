package purse

import (
	"github.com/pkg/errors"
)

type Account struct {
	name string
}

func NewAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return errors.New("account Name is incorrect")
	}

	return nil
}
