package football

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetTeam(t *testing.T) {
	factory := NewTeamFactory()

	teamA1 := factory.GetTeam(TeamA)
	if teamA1 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}
	assert.NotEqual(t, nil, teamA1)

	teamA2 := factory.GetTeam(TeamA)
	if teamA2 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}
	assert.NotEqual(t, nil, teamA2)

	if teamA1 != teamA2 {
		t.Error("TEAM_A objects weren't the same")
	}
	assert.Equal(t, teamA1, teamA2)

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects created was not 1: %d\n", factory.GetNumberOfObjects())
	}
	assert.Equal(t, 1, factory.GetNumberOfObjects())
}

// TestHighVolume is a pressure test
// Create a million calls to the team creation, as a million calls from users.
// Then, it will simply check that the number of teams created is only two.
func TestHighVolume(t *testing.T) {
	factory := NewTeamFactory()

	teams := make([]*Team, 500000*2)
	for i := 0; i < 500000; i++ {
		teams[i] = factory.GetTeam(TeamA)
	}

	for i := 500000; i < 2*500000; i++ {
		teams[i] = factory.GetTeam(TeamB)
	}

	if factory.GetNumberOfObjects() != 2 {
		t.Errorf("The number of objects created was not 2: %d\n", factory.GetNumberOfObjects())
	}
	assert.Equal(t, 2, factory.GetNumberOfObjects())

	// check where the pointers are pointing to, and where they are located. The below is the first three.
	for i := 0; i < 3; i++ {
		fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
	}
}
