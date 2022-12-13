package snake

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

func Start() {
	InitGameObj()
	keyInput := ReadKeyboardInput()
	for !IsGameOver {
		if IsGamePaused {
			DisplayGamePausedInfo()
		}
		ControlSnake(keyInput)
		UpdateGameState()
		displaySnake()
		displayEgg()
		screen.Show()
		time.Sleep(GetSpeedLevel())
	}

	DisplayGameOverInfo()
	time.Sleep(3 * time.Second)
}

// getSnakeHeadCoordinates is used to return a head coordinate.
func getSnakeHeadCoordinates() (int, int) {
	snakeHead := snake.points[len(snake.points)-1]
	return snakeHead.x, snakeHead.y
}

func displaySnake() {
	style := tcell.StyleDefault.Foreground(tcell.ColorLawnGreen.TrueColor())
	for _, snakeCoordinate := range snake.points {
		drawElement(snakeCoordinate.x, snakeCoordinate.y, 1, style, snake.symbol)
	}
}

// getInitialSnakeCoordinates is used to get a snake original coordinate and length.
// If the func has snakeInitialCoordinateLen1 and snakeInitialCoordinateLen2, the snake length is 2.
// If the func has snakeInitialCoordinateLen1, snakeInitialCoordinateLen2, snakeInitialCoordinateLen3, the snake length is 3.
// The coordinates are hard coded to appear somewhere at left section of game frame.
func getInitialSnakeCoordinates() []*Coordinate {
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
// If the snake has eaten an egg, you increase the score and call the function to display the updated score.
// The last thing to check while updating the snake is to determine if the snake’s movement is such that it has bitten itself.
// If it does, the game is over.
func updateSnake() {
	snakeHeadX, snakeHeadY := getSnakeHeadCoordinates()
	newSnakeHead := &Coordinate{
		snakeHeadX + snake.xDirect,
		snakeHeadY + snake.yDirect,
	}
	setSnakeWithinFrame(newSnakeHead)
	// After determining if the snake’s new head coordinate we just need to append to existing coordinates.
	snake.points = append(snake.points, newSnakeHead)

	if !isEggInsideSnake() {
		coordinatesToClear = append(coordinatesToClear, snake.points[0])
		snake.points = snake.points[1:]
	} else {
		score++
		DisplayGameScore()
	}
	if isSnakeEatingItself() {
		IsGameOver = true
	}
}

// setSnakeWithinFrame is used to set snake’s coordinates within game frame and make the snake can pass a border when touch it.
// It is also used to make sure the snake is moving inside the game frame we have defined.
func setSnakeWithinFrame(snakeCoordinate *Coordinate) {
	originX, originY := getFrameTopLeftCoordinate()

	// Determine game frame left boundary is the same as frame’s top left x coordinate
	leftX := originX
	// Determine game frame top boundary is the same as frame’s top left y coordinate
	topY := originY
	// Determine game frame right boundary is equal to sum of left boundary and frame’s with minus frame’s boundary thickness
	rightX := originX + FrameWidth - 1
	// Determine game frame bottom boundary is equal to sum of top boundary and frame’s height
	bottomY := originY + FrameHeight - 1

	// If snake’s y coordinate is less than or equal to top boundary then set new y coordinate as bottom boundary — 1
	if snakeCoordinate.y <= topY {
		snakeCoordinate.y = bottomY - 1
	}
	// If snake’s y coordinate is greater than or equal to bottom boundary then set new y coordinate as top boundary + 1
	if snakeCoordinate.y >= bottomY {
		snakeCoordinate.y = topY + 1
	}
	// If snake’s x coordinate is greater than or equal to right boundary then set new x coordinate as left boundary + 1
	if snakeCoordinate.x >= rightX {
		snakeCoordinate.x = leftX + 1
	}
	// If snake’s x coordinate is less than or equal to left boundary then set new x coordinate as right boundary — 1
	if snakeCoordinate.x <= leftX {
		snakeCoordinate.x = rightX - 1
	}
}

// isSnakeEatingItself is used to check whether snake body coordinates equal to snake head.
// Only the snake eats itself, the game is over.
func isSnakeEatingItself() bool {
	headX, headY := getSnakeHeadCoordinates()
	for _, snakeCoordinate := range snake.points[:len(snake.points)-1] {
		if headX == snakeCoordinate.x && headY == snakeCoordinate.y {
			return true
		}
	}
	return false
}

func UpdateGameState() {
	if IsGamePaused {
		return
	}
	clearEatenEgg()
	updateSnake()
	updateSpeed()
	updateEgg()
}
