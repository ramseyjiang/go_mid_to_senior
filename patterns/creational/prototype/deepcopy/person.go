package deepcopy

import (
	"bytes"
	"encoding/gob"
)

// Prototype - Copy through serialization

// Person Scenario creating person object for 2 employees of the same company
type Person struct {
	Name        string
	Age         int
	WorkAddress *WorkAddress
	HomeAddress *HomeAddress
}

type WorkAddress struct {
	CompanyName, City, Country string
}

type HomeAddress struct {
	HouseNumber, City, Country string
}

// Clone used binary serialization to encode an object and create a copy/clone
func (p *Person) Clone() *Person {
	person := &Person{}
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	_ = enc.Encode(p)
	dec := gob.NewDecoder(&buf)
	_ = dec.Decode(person)
	return person
}
