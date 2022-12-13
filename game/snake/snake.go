package snake

import (
	"github.com/gdamore/tcell/v2"
)

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

var snake *Snake

// getSnakeHeadCoordinates is used to return a head coordinate.
func (sn *Snake) getSnakeHeadCoordinates() (int, int) {
	snakeHead := snake.points[len(snake.points)-1]
	return snakeHead.x, snakeHead.y
}

func (sn *Snake) displaySnake() {
	style := tcell.StyleDefault.Foreground(tcell.ColorLawnGreen.TrueColor())
	for _, snakeCoordinate := range snake.points {
		drawElement(snakeCoordinate.x, snakeCoordinate.y, 1, style, snake.symbol)
		// screen.SetContent(snakeCoordinate.x, snakeCoordinate.y, snake.symbol, nil, style)
	}
}

// getInitialSnakeCoordinates is used to get a snake original coordinate and length.
// If the func has snakeInitialCoordinateLen1 and snakeInitialCoordinateLen2, the snake length is 2.
// If the func has snakeInitialCoordinateLen1, snakeInitialCoordinateLen2, snakeInitialCoordinateLen3, the snake length is 3.
// The coordinates are hard coded to appear somewhere at left section of game frame.
func (sn *Snake) getInitialSnakeCoordinates() []*Coordinate {
	snakeInitialCoordinateLen1 := &Coordinate{8, 5}
	transformCoordinateInsideFrame(snakeInitialCoordinateLen1)

	snakeInitialCoordinateLen2 := &Coordinate{8, 6}
	transformCoordinateInsideFrame(snakeInitialCoordinateLen2)

	// snakeInitialCoordinateLen3 := &Coordinate{8, 7}
	// transformCoordinateInsideFrame(snakeInitialCoordinateLen3)

	// snakeInitialCoordinateLen4 := &Coordinate{8, 8}
	// transformCoordinateInsideFrame(snakeInitialCoordinateLen4)

	return []*Coordinate{
		{snakeInitialCoordinateLen1.x, snakeInitialCoordinateLen1.y},
		{snakeInitialCoordinateLen2.x, snakeInitialCoordinateLen2.y},
		// {snakeInitialCoordinateLen3.x, snakeInitialCoordinateLen3.y},
		// {snakeInitialCoordinateLen4.x, snakeInitialCoordinateLen4.y},
	}
}

// updateSnake is used to check if game should be over when snake eating itself, otherwise grow the snake when the snake eating an egg.
// If the snake does not eat an egg, you need to remove the first coordinate from the snake’s points to maintain its length.
// If the snake has eaten an egg, you increase the game.Score and call the function to display the updated game.Score.
// The last thing to check while updating the snake is to determine if the snake’s movement is such that it has bitten itself.
// If it does, the game is over.
func (sn *Snake) updateSnake() {
	snakeHeadX, snakeHeadY := sn.getSnakeHeadCoordinates()
	newSnakeHead := &Coordinate{
		snakeHeadX + snake.xDirect,
		snakeHeadY + snake.yDirect,
	}
	sn.setSnakeWithinFrame(newSnakeHead)
	// After determining if the snake’s new head coordinate we just need to append to existing coordinates.
	snake.points = append(snake.points, newSnakeHead)

	if !egg.isEggInsideSnake() {
		coordinatesToClear = append(coordinatesToClear, snake.points[0])
		snake.points = snake.points[1:]
	} else {
		game.Score++
		DisplayGameScore()
	}
	if sn.isSnakeEatingItself() {
		game.IsGameOver = true
	}
}

// setSnakeWithinFrame is used to set snake’s coordinates within game frame and make the snake can pass a border when touch it.
// It is also used to make sure the snake is moving inside the game frame we have defined.
func (sn *Snake) setSnakeWithinFrame(snakeCoordinate *Coordinate) {
	originX, originY := getFrameTopLeftCoordinate()

	// Determine game frame left boundary is the same as frame’s top left x coordinate
	leftX := originX
	// Determine game frame top boundary is the same as frame’s top left y coordinate
	topY := originY
	// Determine game frame right boundary is equal to sum of left boundary and frame’s with minus frame’s boundary thickness
	rightX := originX + FrameWidth - 1
	// Determine game frame bottom boundary is equal to sum of top boundary and frame’s height
	bottomY := originY + FrameHeight - 1

	// Below logic is used to make the snake can through the boundary.
	// If snake’s y coordinate is less than or equal to top boundary then set new y coordinate as bottom boundary—1, through the top
	if snakeCoordinate.y <= topY {
		snakeCoordinate.y = bottomY - 1
	}
	// If snake’s y coordinate is greater than or equal to bottom boundary then set new y coordinate as top boundary+1, through the bottom
	if snakeCoordinate.y >= bottomY {
		snakeCoordinate.y = topY + 1
	}
	// If snake’s x coordinate is greater than or equal to right boundary then set new x coordinate as left boundary+1, through the right
	if snakeCoordinate.x >= rightX {
		snakeCoordinate.x = leftX + 1
	}
	// If snake’s x coordinate is less than or equal to left boundary then set new x coordinate as right boundary—1, through the left
	if snakeCoordinate.x <= leftX {
		snakeCoordinate.x = rightX - 1
	}
}

// isSnakeEatingItself is used to check whether snake body coordinates equal to snake head.
// Only the snake eats itself, the game is over.
func (sn *Snake) isSnakeEatingItself() bool {
	headX, headY := sn.getSnakeHeadCoordinates()
	for _, snakeCoordinate := range snake.points[:len(snake.points)-1] {
		if headX == snakeCoordinate.x && headY == snakeCoordinate.y {
			return true
		}
	}
	return false
}
