package snake

import (
	"time"
)

type Game struct {
	Speed        int
	Score        int
	IsGameOver   bool
	IsGamePaused bool
}

var game = &Game{}

func Start() {
	InitGame()
	keyInput := ReadKeyboardInput()
	for !game.IsGameOver {
		if game.IsGamePaused {
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

// InitGame is used to set initial values for the snake variable and the egg variable
// xDirect and yDirect is used to determine the snake init direction.
// xDirect=1, yDirect=0 means leftward.
// xDirect=0, yDirect=-1 means upward.
// xDirect=0, yDirect=1 means downward.
// xDirect=-1, yDirect=0 means rightward.
func InitGame() {
	InitScreen()
	snake = &Snake{
		points:  snake.getInitialSnakeCoordinates(),
		xDirect: 1,
		yDirect: 0,
		symbol:  SymbolSnake,
	}

	egg = &Egg{
		pos:    egg.getInitialEggCoordinates(),
		symbol: SymbolEgg,
	}

	DisplayFrame()
	DisplayGameScore()
	DisplaySpeedLevel()
}

func UpdateGameState() {
	if game.IsGamePaused {
		return
	}
	egg.clearEatenEgg()
	snake.updateSnake()
	updateSpeed()
	egg.updateEgg()
}
