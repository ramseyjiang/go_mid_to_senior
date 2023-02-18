package fsm

import "testing"

func TestGuessNum(t *testing.T) {
	start := StartState{}
	game := GameContext{
		Next: &start,
	}

	for game.Next.executeState(&game) {
	}
}
