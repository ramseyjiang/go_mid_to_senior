package service

import (
	"testing"

	"github.com/ramseyjiang/go_mid_to_senior/pkgusages/testpkgs/suitetest/repos/mocks"
	"github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestVerification struct {
	A      int // First param
	B      int // Second param
	Result int // Expected outcome
}

func TestExampleService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Text example Suite")
}

var _ = Describe("The Example service - example 3", func() {
	var example ExampleService = &ExampleServiceStruct{}

	Context("Function add two values with tax value from DB", func() {
		repositoryMock := new(mocks.ExampleRepository)
		repositoryMock.On("GetExampleTaxValue", mock.Anything, mock.Anything).Return(7)

		// if you want to check exact parameters then you should replace them with real values
		// repositoryMock.On("GetExampleTaxValue", 1, 2).Return(7)
		example = InitExampleService(repositoryMock)

		When("first value have value 50, second 50 and tax rate is 7% ", func() {
			It("should return sum of two values with tax from DB", func() {
				result := example.AddWithTaxValueFromDB(50, 50)

				Expect(result).Should(BeEquivalentTo(107))
			})
		})
	})
})
