package snake

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

func getKeyInput(keyInput chan string) string {
	var key string
	select {
	case key = <-keyInput:
	default:
		key = ""
	}
	return key
}

// ReadKeyboardInput is used to run a goroutine in a loop and returns a channel string from keyboard input.
func ReadKeyboardInput() chan string {
	keyInput := make(chan string)
	go func() {
		for {
			switch eventKey := screen.PollEvent().(type) {
			case *tcell.EventKey:
				keyInput <- eventKey.Name()
			}
		}
	}()
	return keyInput
}

// ControlSnake is used control the snake game base on the player input.
// p or Enter is used to pause the game.
// q or ESC is used to exit the game.
// up arrow or W is used to go up.
// down arrow or S is used to go down.
// left arrow or A is used to go left.
// right arrow or D is used to go right.
func ControlSnake(keyInput chan string) {
	key := getKeyInput(keyInput)
	switch key {
	case "Rune[q]":
		screen.Fini()
		os.Exit(0)
	case "Esc":
		screen.Fini()
		os.Exit(0)
	case "Rune[p]":
		IsGamePaused = !IsGamePaused
	case "Enter":
		IsGamePaused = !IsGamePaused
	case "Up":
		if !IsGamePaused && snake.yDirect == 0 {
			snake.yDirect = -1
			snake.xDirect = 0
		}
	case "Down":
		if !IsGamePaused && snake.yDirect == 0 {
			snake.yDirect = 1
			snake.xDirect = 0
		}
	case "Left":
		if !IsGamePaused && snake.xDirect == 0 {
			snake.yDirect = 0
			snake.xDirect = -1
		}
	case "Right":
		if !IsGamePaused && snake.xDirect == 0 {
			snake.yDirect = 0
			snake.xDirect = 1
		}
	case "Rune[w]":
		if !IsGamePaused && snake.yDirect == 0 {
			snake.yDirect = -1
			snake.xDirect = 0
		}
	case "Rune[s]":
		if !IsGamePaused && snake.yDirect == 0 {
			snake.yDirect = 1
			snake.xDirect = 0
		}
	case "Rune[a]":
		if !IsGamePaused && snake.xDirect == 0 {
			snake.yDirect = 0
			snake.xDirect = -1
		}
	case "Rune[d]":
		if !IsGamePaused && snake.xDirect == 0 {
			snake.yDirect = 0
			snake.xDirect = 1
		}
	}
}
