package deepcopy

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPersonDeepCopy(t *testing.T) {
	person1 := &Person{
		Name: "John",
		Age:  30,
		WorkAddress: &WorkAddress{
			CompanyName: "abc",
			City:        "London",
			Country:     "UK",
		},
		HomeAddress: &HomeAddress{
			HouseNumber: "1A",
			City:        "London",
			Country:     "UK",
		},
	}

	// Clone using serialization, clone person1 object into person2 object.
	person2 := person1.Clone()

	// Modifying name and house address
	person2.Name = "Surya"
	person2.HomeAddress.HouseNumber = "134C"
	person2.HomeAddress.City = "Bangalore"
	person2.HomeAddress.Country = "India"

	assert.Equal(t, "1A", person1.HomeAddress.HouseNumber)
	assert.Equal(t, "London", person1.HomeAddress.City)
	assert.Equal(t, "UK", person1.WorkAddress.Country)
	assert.Equal(t, "134C", person2.HomeAddress.HouseNumber)
	assert.Equal(t, "Bangalore", person2.HomeAddress.City)
	assert.Equal(t, "India", person2.HomeAddress.Country)
}
