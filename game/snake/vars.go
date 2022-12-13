package snake

import "github.com/gdamore/tcell/v2"

// Coordinate will hold multiple integers name x and y representing the position on the terminal screen.
type Coordinate struct {
	x, y int
}

// Snake type consists of 4 things.
// It can move up, down, left, or right depending on user control, and we need to keep it that way unless the user wishes to change the direction.
// A field name points which will be a slice of pointers of type Coordinate, it includes multiple points to represent a snake.
// xDirect is determined the horizontal direction for the snake.
// The yDirect is determined the vertical direction for the snake.
// The symbol is a tag to represent the snake body.
type Snake struct {
	points           []*Coordinate
	xDirect, yDirect int
	symbol           int32
}

// Egg type consists of 2 things.
// point is used to represent the position in the terminal.
// symbol is used to represent an egg's display in the terminal.
type Egg struct {
	point  *Coordinate
	symbol int32
}

var snake *Snake
var egg *Egg

// coordinatesToClear is used to keep track of coordinates to clear after each movement.
var coordinatesToClear []*Coordinate

// screen is used to hold screen information that comes from tcell package
var screen tcell.Screen

// screenWidth and screenHeight are used to hold screen's full width and height respectively.
var screenWidth, screenHeight int

var score, speed int
var (
	IsGameOver   bool
	IsGamePaused bool
)
var err error
