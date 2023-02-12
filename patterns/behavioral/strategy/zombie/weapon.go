package zombie

type Protector struct {
	Name   string
	Weapon Weapon
}

// Weapon is the abstract interface, it has UseWeapon() method only. The client can use that method, don't care how the weapon is triggered.
type Weapon interface {
	UseWeapon() string
}

type SwordWeapon struct {
	Length int
}

func (s SwordWeapon) UseWeapon() string {
	// define slaying action
	return "Slay with a sword"
}

type GunWeapon struct {
	RangeMM int
}

func (gn GunWeapon) UseWeapon() string {
	// define fire action
	return "Fire with a gun"
}

type GrenadeWeapon struct {
	ImpactRadius int
}

func (gr GrenadeWeapon) UseWeapon() string {
	// define throw action
	return "Throw a grenade"
}
