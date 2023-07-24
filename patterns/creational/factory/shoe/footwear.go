package shoe

type Footwear interface {
	SetCategory(name string)
	GetCategory() string

	SetPrice(discount float32)
	GetPrice() float32

	SetGender(gender string)
	GetGender() string
}

type Shoe struct {
	Size     int64
	Price    float32
	Category string
	Gender   string
}

func (s *Shoe) SetCategory(name string) {
	s.Category = name
}

func (s *Shoe) GetCategory() string {
	return s.Category
}

func (s *Shoe) SetPrice(discount float32) {
	s.Price = s.Price * (1 - discount)
}

func (s *Shoe) GetPrice() float32 {
	return s.Price
}

func (s *Shoe) SetGender(gender string) {
	s.Gender = gender
}

func (s *Shoe) GetGender() string {
	return s.Gender
}

func NewShoe(category string, size int64, price float32, gender string, discount float32) Footwear {
	return &Shoe{
		Size:     size,
		Category: category,
		Price:    (1 - discount) * price,
		Gender:   gender,
	}
}
