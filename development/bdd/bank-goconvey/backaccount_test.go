package bankaccounts

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBankAccount(t *testing.T) {
	Convey("Given a user with a valid bank account and sufficient balance", t, func() {
		account1 := NewBankAccount(500.0)
		account2 := NewBankAccount(0.0)

		Convey("When the user transfers money to another valid bank account", func() {
			err := account1.TransferTo(200.0, account2)

			Convey("Then the money should be deducted from their account", func() {
				So(err, ShouldBeNil)
				So(account1.Balance(), ShouldEqual, 300.0)

				Convey("And the money should be received in the other bank account", func() {
					So(account2.Balance(), ShouldEqual, 200.0)
				})
			})
		})
	})
}
