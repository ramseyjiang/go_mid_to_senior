//go:build wireinject
// +build wireinject

package shapes

import "github.com/google/wire"

func ProvideSquare() *Square {
	sq := Square{4}
	return &sq
}

var ShapeSet = wire.NewSet(
	ProvideSquare,
	wire.Bind(new(Geometry), new(*Square)),
)

func ProvideArea(g Geometry) float64 {
	return g.Area()
}

// Because func at here, so it can run "wire" directly in this folder, it can generate wire_gen.go inside.
// Meanwhile, ProvideShape is invoked by wire.go outside this folder, so when it runs "wire", it also can generate
// wire_gen.go outside.

func ProvideShape() float64 {
	panic(wire.Build(ShapeSet, ProvideArea))
	return 1
}
