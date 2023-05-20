package booking

type Booking struct {
	// Fields related to booking
	// (city, date range, hotel room, user info, payment info, etc.)
	// You'd typically have complex types or interfaces here
	// but for simplicity let's use basic types.
	city          string
	dateRange     string
	hotelRoom     string
	personalInfo  string
	paymentInfo   string
	isConfirmed   bool
	receivedEmail bool
}

func (b *Booking) SelectedAValidCityAndDateRange() error {
	// Implement logic for selecting a valid city and date range
	// Here we just hardcode some values
	b.city = "San Francisco"
	b.dateRange = "2023-07-10 to 2023-07-15"

	return nil
}

func (b *Booking) SelectASpecificHotelRoom() error {
	// Implement logic for selecting a specific hotel room
	// Here we just hardcode some value
	b.hotelRoom = "Deluxe Suite"

	return nil
}

func (b *Booking) InputValidPersonalAndPaymentInformation() error {
	// Implement logic for inputting personal and payment information
	// Here we just hardcode some values
	b.personalInfo = "John Doe"
	b.paymentInfo = "Valid Credit Card"

	return nil
}

func (b *Booking) BookingIsConfirmed() error {
	// Implement logic for confirming booking
	// For simplicity, let's just set the field to true
	b.isConfirmed = true

	return nil
}

func (b *Booking) ReceiveABookingConfirmationEmail() error {
	// Implement logic for receiving a booking confirmation email
	// For simplicity, let's just set the field to true
	b.receivedEmail = true

	return nil
}
