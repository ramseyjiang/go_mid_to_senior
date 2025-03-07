package payway

import "fmt"

// an abstract class PaymentProcessor with a ProcessPayment() function that defines the steps of the payment processing algorithm.
// Two concrete classes, CreditCardProcessor and PayPalProcessor, are defined and implement the ProcessPayment() function with specific
// behaviors for processing credit card and PayPal payments, respectively.
// an abstract class PaymentOrder with a ProcessOrder() function that defines the steps of the order processing algorithm.
// A concrete class OnlineOrder is defined and implements the ProcessOrder() function with a reference to a PaymentProcessor object
// that is used to process the payment.

// PaymentProcessor is the Abstract class
type PaymentProcessor interface {
	ProcessPayment(float64) bool
}

// CreditCardProcessor is the Concrete struct 1
type CreditCardProcessor struct{}

// ProcessPayment is an implement for the method of PaymentProcessor.
func (c *CreditCardProcessor) ProcessPayment(amount float64) bool {
	fmt.Println("Processing credit card payment...")
	// code to process credit card payment goes here
	return true
}

// PayPalProcessor is the Concrete struct 2
type PayPalProcessor struct{}

// ProcessPayment is an implement for the method of PaymentProcessor.
func (p *PayPalProcessor) ProcessPayment(amount float64) bool {
	fmt.Println("Processing PayPal payment...")
	// code to process PayPal payment goes here
	return true
}

// PaymentOrder Abstract struct
type PaymentOrder interface {
	ProcessOrder(float64) bool
}

// OnlineOrder Concrete struct
type OnlineOrder struct {
	paymentProcessor PaymentProcessor
}

func (o *OnlineOrder) ProcessOrder(amount float64) bool {
	// Use payment processor to process payment
	paymentProcessed := o.paymentProcessor.ProcessPayment(amount)
	return paymentProcessed
}
