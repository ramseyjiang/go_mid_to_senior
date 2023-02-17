package shop

import "fmt"

// INTERFACES --------------------------------------------------------------

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type Visitor interface {
	Visit(ProductInfoRetriever)
}

type Visitable interface {
	Accept(Visitor)
}

// Product type, implemented the ProductInfoRetriever and the Visitable interfaces, and embedded it on each product.
// Notice: one struct implements two interface together.
type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}

func (p *Product) GetName() string {
	return p.Name
}

// PRODUCTS ----------------------------------------------------------------

type Rice struct {
	Product
}

type Pasta struct {
	Product
}

type Fridge struct {
	Product
}

// GetPrice overrides GetPrice method of Product type
func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

// Accept overrides "Accept" method from Product and implements the Visitable interface
func (f *Fridge) Accept(v Visitor) {
	v.Visit(f)
}

// VISITOR -----------------------------------------------------------------

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NameVisitor struct {
	ProductList string
}

func (n *NameVisitor) Visit(p ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}
