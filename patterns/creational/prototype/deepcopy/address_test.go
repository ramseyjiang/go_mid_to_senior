package deepcopy

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var employee1 = Employee{
	Name:   "John",
	Gender: "Male",
	Address: &Address{
		StreetName: "Baker Street",
		City:       "London",
	},
}

func TestAddressNoneClone(t *testing.T) {
	employee2 := employee1
	employee2.Name = "Surya"
	employee2.Address.StreetName = "Marine Drive"
	employee2.Address.City = "Mumbai"

	// The address is a reference type so even when we create a new employee2 we still are referencing the same Address instance of employee1
	assert.NotEqual(t, "Baker Street", employee1.Address.StreetName)
	assert.NotEqual(t, "London", employee1.Address.City)
}

func TestAddressClone(t *testing.T) {
	employee3 := employee1
	employee3.Name = "Surya"
	employee3.Address = employee1.Address.Clone()
	employee3.Address.StreetName = "Marine Drive"
	employee3.Address.City = "Mumbai"

	assert.Equal(t, "Baker Street", employee1.Address.StreetName)
	assert.Equal(t, "London", employee1.Address.City)
	assert.Equal(t, "Marine Drive", employee3.Address.StreetName)
	assert.Equal(t, "Mumbai", employee3.Address.City)
}
