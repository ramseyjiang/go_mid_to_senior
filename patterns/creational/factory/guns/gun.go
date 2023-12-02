package guns

import (
	"fmt"

	"github.com/pkg/errors"
)

type GunInfo interface {
	getName() string
	getPower() int
}

const Ak = "Ak47"
const Mk = "musket"

type Ak47 struct {
	Name  string
	Power int
}

func (a *Ak47) getName() string {
	return a.Name
}

func (a *Ak47) getPower() int {
	return a.Power
}

type Musket struct {
	Name  string
	Power int
}

func (m *Musket) getName() string {
	return m.Name
}

func (m *Musket) getPower() int {
	return m.Power
}

type GunFactory interface {
	GetGun(gunType string) (GunInfo, error)
}

type GunCreator struct{}

func GetGun(gunType string) (GunInfo, error) {
	switch gunType {
	case Ak:
		return &Ak47{Name: Ak, Power: 5}, nil
	case Mk:
		return &Musket{Name: Mk, Power: 1}, nil
	default:
		return nil, errors.New(fmt.Sprintf("wrong gun type %s passed", gunType))
	}
}
