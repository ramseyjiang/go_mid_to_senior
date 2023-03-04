package purse

import (
	"strconv"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFacade(t *testing.T) {
	// implements the client that uses the Facade to interact with the subsystem
	walletFacade := NewWalletFacade("abc", 1234)
	t.Run("test AddMoneyToWallet", func(t *testing.T) {
		// Implement the client that uses the Facade to interact with the subsystem.
		err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
		if err != nil {
			t.Fatalf("Error: %s\n", err.Error())
		}
		assert.Equal(t, nil, err)
		assert.Equal(t, walletFacade.wallet.balance, 10)
		assert.Equal(t, walletFacade.notification.msg, "Sending wallet credit notification")

		record := "Make ledger entry for accountId" + "abc" + "with txnType " + "credit" + " for amount " + strconv.Itoa(10)
		assert.Equal(t, walletFacade.ledger.record, record)
	})

	t.Run("test DeductMoneyFromWallet", func(t *testing.T) {
		// Implement the client that uses the Facade to interact with the subsystem.
		err := walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
		if err != nil {
			t.Fatalf("Error: %s\n", err.Error())
		}
		assert.Equal(t, nil, err)
		assert.Equal(t, walletFacade.wallet.balance, 5)
		assert.Equal(t, walletFacade.notification.msg, "Sending wallet debit notification")
		record := "Make ledger entry for accountId" + "abc" + "with txnType " + "credit" + " for amount " + strconv.Itoa(5)
		assert.Equal(t, walletFacade.ledger.record, record)
	})
}
