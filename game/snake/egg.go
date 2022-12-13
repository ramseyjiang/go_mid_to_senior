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

func (e *Egg) getInitialEggCoordinates() *Coordinate {
	eggInitialCoordinate := &Coordinate{FrameWidth / 2, FrameHeight / 2}
	transformCoordinateInsideFrame(eggInitialCoordinate)

	return eggInitialCoordinate
}

func (e *Egg) displayEgg() {
	style := tcell.StyleDefault.Foreground(tcell.ColorOrange.TrueColor())
	drawElement(egg.pos.x, egg.pos.y, 1, style, egg.symbol)
}

func (e *Egg) getNewEggCoordinate() (int, int) {
	rand.Seed(time.Now().UnixMicro())
	randomX := rand.Intn(FrameWidth - 3)
	randomY := rand.Intn(FrameHeight - 2)

	newCoordinate := &Coordinate{
		randomX, randomY,
	}

	transformCoordinateInsideFrame(newCoordinate)

	return newCoordinate.x, newCoordinate.y
}

// isEggInsideSnake is used to check whether an agg is eaten by the snake.
// If the egg is eaten, the egg coordinates equal to one of snake body coordinates.
// Only the egg is inside a snake, the snake will grow up.
func (e *Egg) isEggInsideSnake() bool {
	for _, snakeCoordinate := range snake.points {
		if snakeCoordinate.x == egg.pos.x && snakeCoordinate.y == egg.pos.y {
			return true
		}
	}
	return false
}

func (e *Egg) updateEgg() {
	for e.isEggInsideSnake() {
		coordinatesToClear = append(coordinatesToClear, egg.pos)
		egg.pos.x, egg.pos.y = e.getNewEggCoordinate()
	}
}

// clearEatenEgg is used to clear necessary coordinates represented by variable coordinatesToClear
func (e *Egg) clearEatenEgg() {
	for _, coordinate := range coordinatesToClear {
		drawElement(coordinate.x, coordinate.y, 1, tcell.StyleDefault, ' ')
	}
}
