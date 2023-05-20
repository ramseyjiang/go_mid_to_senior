package booking

import (
	"testing"

	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	booking := &Booking{}

	ctx.Step(`^the user has selected a valid city and date range$`, booking.SelectedAValidCityAndDateRange)
	ctx.Step(`^they select a specific hotel room$`, booking.SelectASpecificHotelRoom)
	ctx.Step(`^they input valid personal and payment information$`, booking.InputValidPersonalAndPaymentInformation)
	ctx.Step(`^their booking is confirmed$`, booking.BookingIsConfirmed)
	ctx.Step(`^they receive a booking confirmation email$`, booking.ReceiveABookingConfirmationEmail)
}

func Test_Booking(t *testing.T) {
	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: nil,
		ScenarioInitializer:  InitializeScenario,
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     []string{"features"},
			Randomize: 0,
		},
	}.Run()

	if status != 0 {
		t.Errorf("Some scenarios did not pass")
	}
}
