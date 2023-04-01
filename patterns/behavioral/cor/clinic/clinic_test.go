package clinic

import (
	"log"
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestClinic(t *testing.T) {
	handlerChain := createHandlerChain()

	// patient1 visiting
	patient1 := &Patient{name: "abc"}
	expected1 := []string{"Pharmacy giving medicine to patient", "Reception getting money from patient", "Doctor checking patient", "Reception appointment patient"}
	wanted1 := handlerChain.execute(patient1)
	showStackElement(wanted1)
	assert.Equal(t, true, reflect.DeepEqual(wanted1, expected1))

	// patient2 visiting
	patient2 := &Patient{name: "abc", bookDone: true, pharmacyDone: true}
	expected2 := []string{"Pharmacy already given to patient", "Reception getting money from patient", "Doctor checking patient", "Patient appointment already done"}
	wanted2 := handlerChain.execute(patient2)
	showStackElement(wanted2)
	assert.Equal(t, true, reflect.DeepEqual(wanted2, expected2))

	// patient3 visiting
	patient3 := &Patient{name: "abc", paymentDone: true}
	expected3 := []string{"Pharmacy giving medicine to patient", "Payment Done", "Doctor checking patient", "Reception appointment patient"}
	wanted3 := handlerChain.execute(patient3)
	showStackElement(wanted3)
	assert.Equal(t, true, reflect.DeepEqual(wanted3, expected3))
}

func showStackElement(stack []string) {
	for i := len(stack) - 1; i >= 0; i-- {
		log.Println(stack[i])
	}
	log.Println()
}
