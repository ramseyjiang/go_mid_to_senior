package bookshop

// Book is an interface that specifies the methods that all decorators will implement.
type Book interface {
	GetPrice() float64
}

// BookImpl is a struct that represents the base object that will be decorated.
type BookImpl struct {
	Title string
	Price float64
}

// GetPrice implements the methods for the base object.
func (b *BookImpl) GetPrice() float64 {
	return b.Price
}

// DiscountedBook is a decorator interface that extends the Book interface, and any methods that it should have.
type DiscountedBook interface {
	Book
	SetDiscount(discount float64)
}

// DiscountedBookImpl is a concrete implementation of the decorator interface,
// representing the behavior that will be added to the object being decorated.
type DiscountedBookImpl struct {
	Book     Book
	Discount float64
}

func (d *DiscountedBookImpl) GetPrice() float64 {
	return d.Book.GetPrice() * (1.0 - d.Discount)
}

func (d *DiscountedBookImpl) SetDiscount(discount float64) {
	d.Discount = discount
}
