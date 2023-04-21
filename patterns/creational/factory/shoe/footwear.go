package shoe

type Footwear interface {
	setCategory(name string)
	getCategory() string

	setPrice(discount float32) float32
	getPrice() float32

	setGender(gender string)
	getGender() string
}

type Shoe struct {
	Size     int64
	Price    float32
	Category string
	Gender   string
}

func (s *Shoe) setCategory(name string) {
	s.Category = name
}

func (s *Shoe) getCategory() string {
	return s.Category
}

func (s *Shoe) setPrice(discount float32) float32 {
	return s.Price * (1 - discount)
}

func (s *Shoe) getPrice() float32 {
	return s.Price
}

func (s *Shoe) setGender(gender string) {
	s.Gender = gender
}

func (s *Shoe) getGender() string {
	return s.Gender
}

func newShoe(category string, size int64, price float32, gender string, discount float32) Footwear {
	return &Shoe{
		Size:     size,
		Category: category,
		Price:    (1 - discount) * price,
		Gender:   gender,
	}
}
