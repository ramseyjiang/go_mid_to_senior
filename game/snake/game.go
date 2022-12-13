package snake

import (
	"time"
)

// type Game struct {
// 	snake Snake
// 	egg  Egg
// 	screen tcell.Screen
// }
//
// game := Game{
// 	Snake: snake,
// 	Egg: egg,
// 	Screen: screen,
// }

func Start() {
	InitGameObj()
	keyInput := ReadKeyboardInput()
	for !IsGameOver {
		if IsGamePaused {
			DisplayGamePausedInfo()
		}
		ControlSnake(keyInput)
		UpdateGameState()
		snake.displaySnake()
		egg.displayEgg()
		screen.Show()
		time.Sleep(GetSpeedLevel())
	}

	DisplayGameOverInfo()
	time.Sleep(3 * time.Second)
}

func UpdateGameState() {
	if IsGamePaused {
		return
	}
	egg.clearEatenEgg()
	snake.updateSnake()
	updateSpeed()
	egg.updateEgg()
}
