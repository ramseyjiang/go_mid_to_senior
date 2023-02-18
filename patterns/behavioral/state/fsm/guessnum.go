package fsm

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const target = 100
const retries = 10

// GameState is the state abstract interface
type GameState interface {
	// executeState return true is used to tell the engine that the game must continue. return false is used to tell the game is terminated.
	executeState(*GameContext) bool
}

type GameContext struct {
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState
}

// StartState struct implements the GameState structure because it has the executeState(*GameContext) method of Bool type on its structure.
type StartState struct{}

func (s *StartState) executeState(c *GameContext) bool {
	// it sets the only state possible after executing AskState state.
	c.Next = &AskState{}

	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(target)
	fmt.Println("Introduce a number a number of retries to set the difficulty:")
	c.Retries = retries
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)

	return true
}

type FinishState struct{}

func (f *FinishState) executeState(c *GameContext) bool {
	if c.Won {
		c.Next = &WinState{}
	} else {
		c.Next = &LoseState{}
	}

	return true
}

// AskState struct is the state where we ask the player for a number to guess.
type AskState struct{}

func (a *AskState) executeState(c *GameContext) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n", c.Retries)

	// The test can auto generate a number between 0 to target using the below 2 lines.
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(target)

	fmt.Fscanf(os.Stdin, "%d", &n)
	c.Retries = c.Retries - 1

	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState{}
	}

	if c.Retries == 0 {
		c.Next = &FinishState{}
	}

	return true
}

type WinState struct{}

func (w *WinState) executeState(c *GameContext) bool {
	println("Congrats, you won")

	return false
}

type LoseState struct{}

func (l *LoseState) executeState(c *GameContext) bool {
	fmt.Printf("You loose. The correct number was: %d\n", c.SecretNumber)
	return false
}
