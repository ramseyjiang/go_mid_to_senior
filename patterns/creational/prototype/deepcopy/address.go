package deepcopy

type Employee struct {
	Name    string
	Gender  string
	Address *Address
}

type Address struct {
	StreetName string
	City       string
}

func (add *Address) Clone() *Address {
	return &Address{
		StreetName: add.StreetName,
		City:       add.City,
	}
}
