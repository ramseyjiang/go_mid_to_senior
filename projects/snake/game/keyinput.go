package trisnake

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
	tb "github.com/nsf/termbox-go"
)

var counterSnake = 10
var counterArena = 10

const colorObjectSnake = "Snake"
const colorObjectHard = "Hard"
const colorObjectFood = "Food"
const colorObjectArena = "Arena"

const difficultyEasy = "Easy"
const difficultyNormal = "Normal"

// Tick listens for a keypress and then returns a direction for the snake.
func (snake *Snake) Tick(event tl.Event) {
	// Checks if the event is a keyevent.
	if event.Type == tl.EventKey {
		switch event.Key {
		// Checks if the key is a → press.
		case tl.KeyArrowRight:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != left {
				// Changes the direction of the snake to right.
				snake.Direction = right
			}
		// Checks if the key is a ← press.
		case tl.KeyArrowLeft:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != right {
				// Changes the direction of the snake to left.
				snake.Direction = left
			}
		// Checks if the key is a ↑ press.
		case tl.KeyArrowUp:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != down {
				// Changes the direction of the snake to down.
				snake.Direction = up
			}
		// Checks if the key is a ↓ press.
		case tl.KeyArrowDown:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != up {
				// Changes the direction of the snake to up.
				snake.Direction = down
			}
		}
	}
}

// Tick is a method for the gameoverscreen which listens for either a restart or a quit input from the user.
func (gos *Gameoverscreen) Tick(event tl.Event) {
	// Check if the event is a key event.
	if event.Type == tl.EventKey {
		switch event.Key {
		// If the key pressed is backspace the game will restart!!
		case tl.KeyHome:
			// Will call the RestartGame function to restart the game.
			RestartGame()
		case tl.KeyDelete:
			// Will end the game using a fatal log. This uses the termbox package as termloop does not have a function like that.
			tb.Close()
		case tl.KeySpace:
			SaveHighScore(gs.Score, gs.FPS, Difficulty)
		}
	}
}

// Tick will listen for a keypress to initiate the game.
func (ts *Titlescreen) Tick(event tl.Event) {
	// Checks if the event is a keypress event and the key pressed is the enter key.
	if event.Type == tl.EventKey {
		if event.Key == tl.KeyEnter {
			gs = NewGamescreen()
			sg.Screen().SetLevel(gs)
		}
		if event.Key == tl.KeyInsert {
			gop1 := NewOptionsscreen()
			sg.Screen().SetLevel(gop1)
		}
	}
}

// Tick will listen for a keypress to initiate the game.
func (g *Gameoptionsscreen) Tick(event tl.Event) {
	// Checks if the event is a keypress event.
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyF1:
			ts.GameDifficulty = easy
			Difficulty = difficultyEasy
			gop.CurrentDifficultyText.SetText(fmt.Sprintf("Current difficulty: %s", Difficulty))

		case tl.KeyF2:
			ts.GameDifficulty = normal
			Difficulty = difficultyNormal
			gop.CurrentDifficultyText.SetText(fmt.Sprintf("Current difficulty: %s", Difficulty))

		case tl.KeyArrowUp:
			switch ColorObject {
			case colorObjectSnake:
				if counterSnake <= 10 {
					return
				}
				counterSnake -= 2
				gop.ColorSelectedIcon.SetPosition(73, counterSnake)
			case colorObjectArena:
				if counterArena <= 10 {
					return
				}
				counterArena -= 2
				gop.ColorSelectedIcon.SetPosition(73, counterArena)
			}
		case tl.KeyArrowDown:
			switch ColorObject {
			case colorObjectSnake:
				if counterSnake >= 22 {
					return
				}
				counterSnake += 2
				gop.ColorSelectedIcon.SetPosition(73, counterSnake)

			case colorObjectArena:
				if counterArena >= 22 {
					return
				}
				counterArena += 2
				gop.ColorSelectedIcon.SetPosition(73, counterArena)

			}

		case tl.KeyF3:
			ts.GameDifficulty = hard
			Difficulty = colorObjectHard
			gop.CurrentDifficultyText.SetText(fmt.Sprintf("Current difficulty: %s", Difficulty))

		case tl.KeyF4:
			ColorObject = colorObjectSnake
			gop.CurrentColorObjectText.SetText(fmt.Sprintf("Current object: %s", ColorObject))

		case tl.KeyF5:
			ColorObject = colorObjectFood
			gop.CurrentColorObjectText.SetText(fmt.Sprintf("Current object: %s", ColorObject))

		case tl.KeyF6:
			ColorObject = colorObjectArena
			gop.CurrentColorObjectText.SetText(fmt.Sprintf("Current object: %s", ColorObject))

		case tl.KeyEnter:
			gs = NewGamescreen()
			sg.Screen().SetLevel(gs)
		}
	}
}

func CheckSelectedColor(c int) tl.Attr {
	switch c {
	case white:
		return tl.ColorWhite
	case red:
		return tl.ColorRed
	case green:
		return tl.ColorGreen
	case blue:
		return tl.ColorBlue
	case yellow:
		return tl.ColorYellow
	case magenta:
		return tl.ColorMagenta
	case cyan:
		return tl.ColorCyan
	default:
		return tl.ColorDefault
	}
}
