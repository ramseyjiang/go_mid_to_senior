package shop

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/google/uuid"
)

func TestClone(t *testing.T) {
	shirtColors := GetShirtsCloner()
	if shirtColors == nil {
		t.Fatal("Received cache was nil")
	}
	assert.Equal(t, shirtColors != nil, true)

	item1, _ := shirtColors.GetClone(White)
	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}
	shirt1.SN = uuid.New().String() // set a new SN, make sure all clones have different SN.

	item2, err := shirtColors.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}
	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt2 couldn't be done successfully")
	}
	shirt2.SN = uuid.New().String() // set a new SN, make sure all clones have different SN.

	if shirt1.SN == shirt2.SN {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}
	assert.NotEqual(t, shirt1.SN, shirt2.SN)
	assert.NotEqual(t, shirt1, shirt2)
	assert.Equal(t, shirt1.Color, shirt2.Color)

	if shirt1 == shirt2 {
		t.Error("Shirt 1 cannot be equal to Shirt 2")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())

	t.Logf("LOG: The memory positions of the shirts are different %p != %p", &shirt1, &shirt2)
}
