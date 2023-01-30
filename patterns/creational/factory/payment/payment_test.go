package payment

import (
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCreatePayWayCash(t *testing.T) {
	payment, err := GetPayWay(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}
	assert.Equal(t, CashName, payment.GetName())

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPayWayDebitCard(t *testing.T) {
	payment, err := GetPayWay(DebitCard)
	if err != nil {
		t.Error("A payment method of type 'DebitCard' must exist")
	}
	assert.Equal(t, DebitName, payment.GetName())

	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPayWayCreditCard(t *testing.T) {
	payment, err := GetPayWay(CreditCard)
	if err != nil {
		t.Error("A payment method of type 'CreditCard' must exist")
	}
	assert.Equal(t, CreditName, payment.GetName())

	msg := payment.Pay(32.30)
	if !strings.Contains(msg, "paid using credit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPayWay(20)
	if err == nil {
		t.Error("A payment method with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}
