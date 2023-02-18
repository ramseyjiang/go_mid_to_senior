package shop

// The ProductInfoRetriever type has a method to get the price and the name of the product.
type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

// The Visitor interface has a Visit method that accepts the ProductInfoRetriever type as a visitor.
type Visitor interface {
	Visit(ProductInfoRetriever)
}

// Visitable interface has an Accept method that takes a Visitor type as an argument.
type Visitable interface {
	Accept(Visitor)
}

// Product type, implemented the ProductInfoRetriever and the Visitable interfaces together, and embedded it on each product.
// The Product type that can store the information about almost any product of the shop.
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

// Rice is one type of the product, it has the Product type embedded.
type Rice struct {
	Product
}

// Pasta is one type of the product, it has the Product type embedded.
type Pasta struct {
	Product
}

// Fridge is one type of the product, it has the Product type embedded.
type Fridge struct {
	Product
}

// GetPrice overrides GetPrice method of Product type
func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 30
}

// Accept overrides "Accept" method from Product and implements the Visitable interface
func (f *Fridge) Accept(v Visitor) {
	v.Visit(f)
}

// a couple of Visitors interfaces are below.

// PriceVisitor sums the price of all products
type PriceVisitor struct {
	Sum float32
}

// The Visit takes the value of the Price variable of the ProductInfoRetriever type, passed as an argument, and adds it to the Sum field.
func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

// NameVisitor prints the name of each product
type NameVisitor struct {
	ProductList []string
}

// Visit stores the name of the ProductInfoRetriever type, passed as an argument, and appends it to a new line on the ProductList field.
func (n *NameVisitor) Visit(p ProductInfoRetriever) {
	n.ProductList = append(n.ProductList, p.GetName())
}
