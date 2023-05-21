package ecommerce

type Product struct {
	Name  string
	Price float64
}

type Cart struct {
	Products []Product
}

func (c *Cart) AddProduct(p Product) {
	c.Products = append(c.Products, p)
}

func (c *Cart) Checkout() {
	c.Products = []Product{} // Empty the cart
}

func (c *Cart) TotalItems() int {
	return len(c.Products)
}
