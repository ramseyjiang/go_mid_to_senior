package repos

type ExampleRepository interface {
	GetExampleTaxValue(firstParm, secondParam int) int
}

func (repo *ExampleRepositoryStruct) GetExampleTaxValue(firstParm, secondParam int) int {
	return 0
}

type ExampleRepositoryStruct struct {
}
