package shoe

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSlipper(t *testing.T) {
	slipper, _ := getShoe("Slipper", 41, "Male", 200)
	assert.Equal(t, "Slipper", slipper.getCategory())
	assert.Equal(t, "Male", slipper.getGender())
	assert.Equal(t, float32(200*(1-SlipperDiscount)), slipper.getPrice())
}

func TestSandal(t *testing.T) {
	sandal, _ := getShoe("Sandal", 42, "Male", 1000)
	assert.Equal(t, "Sandal", sandal.getCategory())
	assert.Equal(t, "Male", sandal.getGender())
	assert.Equal(t, float32(1000*(1-SandalDiscount)), sandal.getPrice())
}

func TestBucks(t *testing.T) {
	bucks, _ := getShoe("Bucks", 38, "Female", 100)
	assert.Equal(t, "Bucks", bucks.getCategory())
	assert.Equal(t, "Female", bucks.getGender())
	assert.Equal(t, float32(100*(1-BucksDiscount)), bucks.getPrice())
}

func TestNoneShoeType(t *testing.T) {
	_, err := getShoe("test", 38, "Female", 100)
	if err == nil {
		t.Error("invalid footwear type")
	}
	t.Log("LOG:", err)
}
