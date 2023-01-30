package guns

import (
	"fmt"

	"github.com/pkg/errors"
)

type GunInfo interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

const ak = "ak47"
const mk = "musket"

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

func getGun(gunType string) (GunInfo, error) {
	switch gunType {
	case ak:
		return newAk47(), nil
	case mk:
		return newMusket(), nil
	default:
		return nil, errors.New(fmt.Sprintf("wrong gun type %s passed", gunType))
	}
}
