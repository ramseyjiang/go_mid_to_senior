package shoe

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSlipper(t *testing.T) {
	// Use the Factory Pattern in the client code. Step 5
	factory := NewConcreteShoeFactory()
	slipper, _ := factory.CreateShoe(Slipper, 41, "Male", 200)
	assert.Equal(t, Slipper, slipper.GetCategory())
	assert.Equal(t, "Male", slipper.GetGender())
	assert.Equal(t, float32(200*(1-SlipperDiscount)), slipper.GetPrice())
}

func TestSandal(t *testing.T) {
	// Use the Factory Pattern in the client code. Step 5
	factory := NewConcreteShoeFactory()
	sandal, _ := factory.CreateShoe(Sandal, 42, "Male", 1000)
	assert.Equal(t, Sandal, sandal.GetCategory())
	assert.Equal(t, "Male", sandal.GetGender())
	assert.Equal(t, float32(1000*(1-SandalDiscount)), sandal.GetPrice())
}

func TestBucks(t *testing.T) {
	// Use the Factory Pattern in the client code. Step 5
	factory := NewConcreteShoeFactory()
	bucks, _ := factory.CreateShoe(Bucks, 38, "Female", 100)
	assert.Equal(t, Bucks, bucks.GetCategory())
	assert.Equal(t, "Female", bucks.GetGender())
	assert.Equal(t, float32(100*(1-BucksDiscount)), bucks.GetPrice())
}

func TestNoneShoeType(t *testing.T) {
	// Use the Factory Pattern in the client code. Step 5
	factory := NewConcreteShoeFactory()
	_, err := factory.CreateShoe("test", 38, "Female", 100)
	if err == nil {
		t.Error("invalid footwear type")
	}
	t.Log("LOG:", err)
}
