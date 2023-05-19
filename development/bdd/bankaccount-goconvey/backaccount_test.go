package bankaccount

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBankAccount(t *testing.T) {
	Convey("Given a bank account with sufficient balance", t, func() {
		account := NewBankAccount(500)

		Convey("When the user withdraws a valid amount", func() {
			err := account.Withdraw(200)

			Convey("Then the withdrawal should be successful", func() {
				So(err, ShouldBeNil)

				Convey("And the balance should be reduced", func() {
					So(account.Balance(), ShouldEqual, 300)
				})
			})
		})

		Convey("When the user withdraws an amount greater than the balance", func() {
			err := account.Withdraw(600)

			Convey("Then the withdrawal should fail", func() {
				So(err, ShouldNotBeNil)

				Convey("And the balance should not change", func() {
					So(account.Balance(), ShouldEqual, 500)
				})
			})
		})
	})
}
