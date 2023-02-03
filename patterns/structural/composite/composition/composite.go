package composition

type Athlete struct{}

func (a *Athlete) Train() string {
	return "Training"
}

func Swim() string {
	return "Swimming!"
}

type CompositeSwimmerA struct {
	AthleteA Athlete
	SwimmerA *func() string
}

// --------------------------------------------------------------------------

type Trainer interface {
	Train() string
}
type Swimmer interface {
	Swim() string
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() string {
	return "Swimming!"
}

// CompositeSwimmerB has two embed elements, both of them are point to interface separately.
type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

// --------------------------------------------------------------------------

type Animal struct{}

func (r *Animal) Eat() string {
	return "Eating"
}

type Shark struct {
	Animal
	Swim func() string
}

// --------------------------------------------------------------------------

// Tree is another very common approach always using the Composite pattern.
// This is some kind of recursive compositing, and, because of the nature of recursive,
// we must use pointers so that the compiler knows how much memory it must reserve for this struct.
type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

// --------------------------------------------------------------------------

type Parent struct {
	SomeField int
}

// Son struct is composite with the parent without embedding, then using GetParentField can get a parent field easily.
type Son struct {
	P Parent
}

func GetParentField(p Parent) int {
	return p.SomeField
}
