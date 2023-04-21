package guns

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAK47(t *testing.T) {
	ak47, _ := GetGun(Ak)
	assert.Equal(t, Ak, ak47.getName())
	assert.Equal(t, 5, ak47.getPower())
}

func TestMusket(t *testing.T) {
	mk, _ := GetGun(Mk)
	assert.Equal(t, Mk, mk.getName())
	assert.Equal(t, 1, mk.getPower())
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetGun("sk")
	if err == nil {
		t.Error("A gun type with sk must return an error")
	}
	t.Log("LOG:", err)
}
