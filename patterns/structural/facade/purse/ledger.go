package purse

import (
	"strconv"
)

// Ledger is the complex subsystem's structs.
type Ledger struct {
	record string
}

func (l *Ledger) makeEntry(accountID string, txnType string, amount int) {
	l.record = "Make ledger entry for accountId" + accountID + "with txnType " + txnType + " for amount " + strconv.Itoa(amount)
}
