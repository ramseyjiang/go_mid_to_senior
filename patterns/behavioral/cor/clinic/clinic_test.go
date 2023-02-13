package clinic

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestClinic(t *testing.T) {
	pharmacy := &Pharmacy{}
	payment := &Payment{}
	doctor := &Doctor{}
	appointment := &Appointment{}

	// Set next for doctor step
	payment.setNext(pharmacy)

	// Set next for payment step
	doctor.setNext(payment)

	// Set next for pharmacy step
	appointment.setNext(doctor)

	// patient1 visiting
	patient1 := &Patient{name: "abc"}
	expected1 := "Pharmacy giving medicine to patient, Reception getting money from patient, Doctor checking patient, Reception appointment patient"
	assert.Equal(t, expected1, appointment.execute(patient1))

	// patient2 visiting
	patient2 := &Patient{name: "abc", bookDone: true, pharmacyDone: true}
	excepted2 := "Pharmacy already given to patient, Reception getting money from patient, Doctor checking patient, Patient appointment already done"
	assert.Equal(t, excepted2, appointment.execute(patient2))
}
