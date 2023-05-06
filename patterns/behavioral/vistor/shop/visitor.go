package shop

type Item interface {
	Accept(visitor DiscountVisitor) float64
}

type DiscountVisitor interface {
	VisitBook(book *Book) float64
	VisitElectronic(electronic *Electronic) float64
}

type Book struct {
	price float64
}

func (b *Book) Accept(visitor DiscountVisitor) float64 {
	return visitor.VisitBook(b)
}

type Electronic struct {
	price float64
}

func (e *Electronic) Accept(visitor DiscountVisitor) float64 {
	return visitor.VisitElectronic(e)
}

type SeasonalDiscountVisitor struct {
	bookDiscount       float64
	electronicDiscount float64
}

func (v *SeasonalDiscountVisitor) VisitBook(book *Book) float64 {
	return book.price * (1 - v.bookDiscount)
}

func (v *SeasonalDiscountVisitor) VisitElectronic(electronic *Electronic) float64 {
	return electronic.price * (1 - v.electronicDiscount)
}
