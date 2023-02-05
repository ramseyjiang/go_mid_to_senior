package payway

// ItemPayment is Abstraction interface
type ItemPayment interface {
	Purchase() string
}

type Item struct {
	Name    string
	Price   float64
	Payment PaymentMethod
}

func (i *Item) Purchase() string {
	return i.Payment.Pay(i.Price)
}
