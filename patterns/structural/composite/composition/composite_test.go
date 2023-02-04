package composition

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAthleteTrain(t *testing.T) {
	athlete := Athlete{}
	assert.Equal(t, "Training", athlete.Train())
}

func TestSwimmerASwim(t *testing.T) {
	// As we want to use Swim() function in CompositeSwimmerA struct, we have to make a two-step call.
	// Firstly, assign function Swim() to the localSwim visible, because a function doesn't have an address to pass it to the CompositeSwimmerA type
	// Secondly, take the localSwim and copy it to SwimmerA method
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		SwimmerA: &localSwim,
	}

	assert.Equal(t, "Training", swimmer.AthleteA.Train())
	assert.Equal(t, "Swimming!", (*swimmer.SwimmerA)()) // SwimmerA is a closure in swimmer
}

func TestAnimalSwim(t *testing.T) {
	// Notice the Swim here can use it directly, because SwimmerA of CompositeSwimmerA is defined with a pointer function type.
	// the Swim of Shark is defined only the function type. So it does not need to covert to an address first.
	fish := Shark{
		Swim: Swim,
	}

	// fish.Eat() that is used embed objects.
	// In Golang, you can also embed objects within objects to make it look a lot like inheritance.
	// That is, we won't have to explicitly call the field name to have access to its fields and method because they'll be part of us.
	assert.Equal(t, "Eating", fish.Eat())
	assert.Equal(t, "Swimming!", fish.Swim())
}

func TestSwimmerBSwim(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImplementor{},
	}

	assert.Equal(t, "Training", swimmer.Train())
	assert.Equal(t, "Swimming!", swimmer.Swim())
}

func TestTree(t *testing.T) {
	// This tree struct stored a LeafValue object for each instance and a new Tree in its Right and Left fields.
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
		},
		Left: &Tree{4, nil, nil},
	}

	assert.Equal(t, 5, root.Right.LeafValue)
	assert.Equal(t, 6, root.Right.Right.LeafValue)
	assert.Equal(t, nil, root.Right.Right.Right)
	assert.Equal(t, 4, root.Left.LeafValue)
}

func TestSonGetParentField(t *testing.T) {
	son := Son{}
	assert.Equal(t, 0, GetParentField(son.P)) // In golang, int type default value is 0.
}
