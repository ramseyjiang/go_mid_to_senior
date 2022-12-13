package snake

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Egg type consists of 2 things.
// pos is used to represent the position in the terminal.
// symbol is used to represent an egg's display in the terminal.
type Egg struct {
	pos    *Coordinate
	symbol int32
}

var egg *Egg

func getInitialEggCoordinates() *Coordinate {
	eggInitialCoordinate := &Coordinate{FrameWidth / 2, FrameHeight / 2}
	transformCoordinateInsideFrame(eggInitialCoordinate)

	return eggInitialCoordinate
}

func displayEgg() {
	style := tcell.StyleDefault.Foreground(tcell.ColorOrange.TrueColor())
	drawElement(egg.pos.x, egg.pos.y, 1, style, egg.symbol)
}

func getNewEggCoordinate() (int, int) {
	rand.Seed(time.Now().UnixMicro())
	randomX := rand.Intn(FrameWidth - 3)
	randomY := rand.Intn(FrameHeight - 1)

	newCoordinate := &Coordinate{
		randomX, randomY,
	}

	transformCoordinateInsideFrame(newCoordinate)

	return newCoordinate.x, newCoordinate.y
}

// isEggInsideSnake is used to check whether an agg is eaten by the snake.
// If the egg is eaten, the egg coordinates equal to one of snake body coordinates.
// Only the egg is inside a snake, the snake will grow up.
func isEggInsideSnake() bool {
	for _, snakeCoordinate := range snake.points {
		if snakeCoordinate.x == egg.pos.x && snakeCoordinate.y == egg.pos.y {
			return true
		}
	}
	return false
}

func updateEgg() {
	for isEggInsideSnake() {
		coordinatesToClear = append(coordinatesToClear, egg.pos)
		egg.pos.x, egg.pos.y = getNewEggCoordinate()
	}
}

// clearEatenEgg is used to clear necessary coordinates represented by variable coordinatesToClear
func clearEatenEgg() {
	for _, coordinate := range coordinatesToClear {
		drawElement(coordinate.x, coordinate.y, 1, tcell.StyleDefault, ' ')
	}
}
