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

func (c *Cart) TotalPrice() float64 {
	var total float64
	for _, p := range c.Products {
		total += p.Price
	}
	return total
}
