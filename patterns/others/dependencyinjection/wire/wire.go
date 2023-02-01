//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/ramseyjiang/go_mid_to_senior/patterns/others/dependencyinjection/wire/foobarbaz"
	"github.com/ramseyjiang/go_mid_to_senior/patterns/others/dependencyinjection/wire/greeter"
	"github.com/ramseyjiang/go_mid_to_senior/patterns/others/dependencyinjection/wire/shapes"
)

// In Wire, initializers are known as "providers," functions which provide a particular type.
// We add a zero value for Event as a return value to satisfy the compiler.
// Note that even if we add values to Event, Wire will ignore them.
// In fact, the injector's purpose is to provide information about which providers to use to construct an Event.

// Install the tool with:
// go get github.com/google/wire/cmd/wire
// Then run "% wire" in CLI, it will generate wire_gen.go automatically.

// Wire inspects the arguments to the injector, sees that we added a string to the list of arguments (e.g., phrase),
// and likewise sees that among all the providers, NewMessage takes a string, and so it passes phrase into NewMessage.

// After wire_gen.go is generated, InitializeEvent will in the wire.go and wire_gen.go,
// Don't worry. Each time you update in provider.go and wire.go.
// Then run "% wire" in you CLI, it will regenerate wire_gen.go

// An injector is declared by writing a function declaration whose body is a call to wire.Build.

func InitializeEvent() (greeter.Event, error) {
	wire.Build(greeter.SuperSet)
	return greeter.Event{}, nil
}

// An injector is declared by writing a function declaration whose body is a call to wire.Build.
// The return values don't matter as long as they are of the correct type.

func InitializeBaz() (foobarbaz.Baz, error) {
	wire.Build(foobarbaz.SuperSet)
	return foobarbaz.Baz{}, nil
}

/**
Let's summarize what we have done here. First, we wrote a number of components with corresponding initializers,
or providers. Next, we created an injector function, specifying which arguments it receives and which types it returns.
Then, we filled in the injector function with a call to wire.Build supplying all necessary providers.
Finally, we ran the wire command to generate code that wires up all the different initializers.
When we added an argument to the injector and an error return value,
running wire again made all the necessary updates to our generated code.

The example here is small, but it demonstrates some of the power of Wire,
and how it takes much of the pain out of initializing code using dependency injection.
Furthermore, using Wire produced code that looks much like what we would otherwise write.
There are no bespoke types that commit a user to Wire. Instead it's just generated code.
We may do with it what we will.
Finally, another point worth considering is how easy it is to add new dependencies to our component initialization.
As long as we tell Wire how to provide (i.e., initialize) a component,
we may add that component anywhere in the dependency graph and Wire will handle the rest.
*/

func ProvideShape() float64 {
	panic(wire.Build(shapes.ShapeSet, shapes.ProvideArea))
	return 1
}
