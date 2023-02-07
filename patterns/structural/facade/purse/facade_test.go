package purse

import (
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFacade(t *testing.T) {
	walletFacade := newWalletFacade("abc", 1234)
	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	assert.Equal(t, nil, err)

	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	assert.Equal(t, nil, err)
}
