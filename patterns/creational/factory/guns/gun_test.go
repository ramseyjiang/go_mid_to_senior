package guns

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAK47(t *testing.T) {
	ak47, _ := getGun(ak)
	assert.Equal(t, "AK47 gun", ak47.getName())
	assert.Equal(t, 4, ak47.getPower())
}

func TestMusket(t *testing.T) {
	mk, _ := getGun(mk)
	assert.Equal(t, "Musket gun", mk.getName())
	assert.Equal(t, 1, mk.getPower())
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := getGun("sk")
	if err == nil {
		t.Error("A gun type with sk must return an error")
	}
	t.Log("LOG:", err)
}
