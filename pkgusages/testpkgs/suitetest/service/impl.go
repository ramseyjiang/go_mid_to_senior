package service

import "github.com/ramseyjiang/go_mid_to_senior/pkgusages/testpkgs/suitetest/repos"

type ExampleService interface {
	Add(a, b int) int
	AddWithTaxValueFromDB(a, b int) float32
}

type ExampleServiceStruct struct {
	exampleRepository repos.ExampleRepository
}

func InitExampleService(exampleRepository repos.ExampleRepository) ExampleService {
	service := new(ExampleServiceStruct)
	service.exampleRepository = exampleRepository

	return service
}

func (example *ExampleServiceStruct) AddWithTaxValueFromDB(a, b int) float32 {
	taxVal := example.exampleRepository.GetExampleTaxValue(a, b)
	taxValFloat := float32(taxVal) / 100
	sum := float32(a + b)

	return sum + sum*taxValFloat
}

func (example *ExampleServiceStruct) Add(a, b int) int {
	return a + b
}
