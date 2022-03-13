package factory

type iFootwear interface {
	setCategory(name string)
	getCategory() string

	setPrice(discount float32) float32
	getPrice() float32

	setGender(gender string)
	getGender() string
}

type shoe struct {
	size     int64
	price    float32
	category string
	gender   string
}

func (s *shoe) setCategory(name string) {
	s.category = name
}

func (s *shoe) getCategory() string {
	return s.category
}

func (s *shoe) setPrice(discount float32) float32 {
	return s.price * (1 - discount)
}

func (s *shoe) getPrice() float32 {
	return s.price
}

func (s *shoe) setGender(gender string) {
	s.gender = gender
}

func (s *shoe) getGender() string {
	return s.gender
}

func newShoe(size int64, category string, price float32, gender string, discount float32) iFootwear {
	return &shoe{
		size:     size,
		category: category,
		price:    (1 - discount) * price,
		gender:   gender,
	}
}
