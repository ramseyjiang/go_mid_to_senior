package zombie

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// Weapon.UseWeapon() in order to use the weapon, no matter what weapon it is.
// This design provides us with the flexibility to switch between the various implementations for a single behavior.

func TestProtector(t *testing.T) {
	weapon := GunWeapon{10}
	tom := Protector{Name: "tom", Weapon: weapon}
	tom.Weapon.UseWeapon()
	assert.Equal(t, "Fire with a gun", tom.Weapon.UseWeapon())

	swordWeapon := SwordWeapon{10}
	tom.Weapon = swordWeapon
	assert.Equal(t, "Slay with a sword", tom.Weapon.UseWeapon())

	davy := Protector{Name: "davy", Weapon: GrenadeWeapon{10}}
	assert.Equal(t, "Throw a grenade", davy.Weapon.UseWeapon())
}
