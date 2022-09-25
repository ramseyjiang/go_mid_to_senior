package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

// Coordinate will hold multiple integers name x and y representing the position on the terminal screen.
type Coordinate struct {
	x, y int
}

// Snake type consists of 4 things.
// It can move up, down, left, or right depending on user control, and we need to keep it that way unless the user wishes to change the direction.
// A field name points which will be a slice of pointers of type Coordinate, it includes multiple points to represent a snake.
// horizontalDirect is determined the horizontal direction for the snake.
// The verticalDirect is determined the vertical direction for the snake.
// The symbol is a tag to represent the snake body.
type Snake struct {
	points                           []*Coordinate
	horizontalDirect, verticalDirect int
	symbol                           rune
}

// Egg type consists of 2 things.
// point is used to represent the position in the terminal.
// symbol is used to represent an egg's display in the terminal.
type Egg struct {
	point  *Coordinate
	symbol rune
}

var snake *Snake
var egg *Egg

// coordinatesToClear is used to keep track of coordinates to clear after each movement.
var coordinatesToClear []*Coordinate

// Screen is used to hold Screen information that comes from tcell package
var Screen tcell.Screen

// screenWidth and screenHeight are used to hold Screen's full width and height respectively.
var screenWidth, screenHeight int

var score, speed int
var isGamePaused, isGameOver bool

// SnakeSymbol and EggSymbol is used to display snake and egg position at the beginning.
// 0x2588 means a rectangle, 0x25CF means a cycle point.
const SnakeSymbol = 0x2588
const EggSymbol = 0x25CF

// FrameWidth and FrameHeight is used to change it according to your need
// When run this game, please let your terminal has enough height and width, otherwise, please adjust screenWidth and screenHeight value.
const FrameWidth = 80
const FrameHeight = 20

// FrameBorderThickness means game frame's border thickness
const FrameBorderThickness = 1

// FrameBorderVertical and others const below is used to represent game frame border symbols to make it look a bit fancy
const FrameBorderVertical = '║'
const FrameBorderHorizontal = '═'
const FrameBorderTopLeft = '╔'
const FrameBorderTopRight = '╗'
const FrameBorderBottomRight = '╝'
const FrameBorderBottomLeft = '╚'

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	initScreen()
	initGameObj()
	displayFrame()
	displayGameScore()
	displaySpeedLevel()
	keyInput := readKeyboardInput()
	for !isGameOver {
		if isGamePaused {
			displayGamePausedInfo()
		}
		gameControl(keyInput)
		updateGameState()
		displayGameObjects()
		time.Sleep(getSpeedLevel())
	}

	displayGameOverInfo()
	time.Sleep(3 * time.Second)
}

func getSpeedLevel() time.Duration {
	speed1 := 120 * time.Millisecond
	speed2 := 90 * time.Millisecond
	speed3 := 60 * time.Millisecond
	speed4 := 30 * time.Millisecond
	scoreLevel1 := 1
	scoreLevel2 := 2
	scoreLevel3 := 3

	if score > scoreLevel1 && score <= scoreLevel2 {
		speed = 1
		return speed2
	}

	if score > scoreLevel2 && score <= scoreLevel3 {
		speed = 2
		return speed3
	}

	if score > scoreLevel3 {
		speed = 3
		return speed4
	}

	return speed1
}

func getKeyInput(keyInput chan string) string {
	var key string
	select {
	case key = <-keyInput:
	default:
		key = ""
	}
	return key
}

// readKeyboardInput is used to run a goroutine in a loop and returns a channel string from keyboard input.
func readKeyboardInput() chan string {
	keyInput := make(chan string)
	go func() {
		for {
			switch ev := Screen.PollEvent().(type) {
			case *tcell.EventKey:
				keyInput <- ev.Name()
			}
		}
	}()
	return keyInput
}

// gameControl is used control the snake game base on the player input.
// p or Enter is used to pause the game.
// q or ESC is used to exit the game.
// up arrow or w is used to go up.
// down arrow or s is used to go down.
// left arrow or a is used to go left.
// right arrow or w is used to go right.
func gameControl(keyInput chan string) {
	key := getKeyInput(keyInput)
	switch key {
	case "Rune[q]":
		Screen.Fini()
		os.Exit(0)
	case "Esc":
		Screen.Fini()
		os.Exit(0)
	case "Rune[p]":
		isGamePaused = !isGamePaused
	case "Enter":
		isGamePaused = !isGamePaused
	case "Up":
		if !isGamePaused && snake.verticalDirect == 0 {
			snake.verticalDirect = -1
			snake.horizontalDirect = 0
		}
	case "Down":
		if !isGamePaused && snake.verticalDirect == 0 {
			snake.verticalDirect = 1
			snake.horizontalDirect = 0
		}
	case "Left":
		if !isGamePaused && snake.horizontalDirect == 0 {
			snake.verticalDirect = 0
			snake.horizontalDirect = -1
		}
	case "Right":
		if !isGamePaused && snake.horizontalDirect == 0 {
			snake.verticalDirect = 0
			snake.horizontalDirect = 1
		}
	case "Rune[w]":
		if !isGamePaused && snake.verticalDirect == 0 {
			snake.verticalDirect = -1
			snake.horizontalDirect = 0
		}
	case "Rune[s]":
		if !isGamePaused && snake.verticalDirect == 0 {
			snake.verticalDirect = 1
			snake.horizontalDirect = 0
		}
	case "Rune[a]":
		if !isGamePaused && snake.horizontalDirect == 0 {
			snake.verticalDirect = 0
			snake.horizontalDirect = -1
		}
	case "Rune[d]":
		if !isGamePaused && snake.horizontalDirect == 0 {
			snake.verticalDirect = 0
			snake.horizontalDirect = 1
		}
	}
}

// getSnakeHeadCoordinates is used to return a head coordinate.
func getSnakeHeadCoordinates() (int, int) {
	snakeHead := snake.points[len(snake.points)-1]
	return snakeHead.x, snakeHead.y
}

// updateSnake is used to check if game should be over when snake eating itself, otherwise grow the snake when the snake eating an egg.
// If the snake does not eat an egg, you need to remove the first coordinate from the snake’s points to maintain its length.
// If the snake has eaten an egg, you increase the score and call the function to display the updated score.
// The last thing to check while updating the snake is to determine if the snake’s movement is such that it has bitten itself.
// If it does, the game is over.
func updateSnake() {
	snakeHeadX, snakeHeadY := getSnakeHeadCoordinates()
	newSnakeHead := &Coordinate{
		snakeHeadX + snake.horizontalDirect,
		snakeHeadY + snake.verticalDirect,
	}
	setSnakeWithinFrame(newSnakeHead)
	// After determining if the snake’s new head coordinate we just need to append to existing coordinates.
	snake.points = append(snake.points, newSnakeHead)

	if !isEggInsideSnake() {
		coordinatesToClear = append(coordinatesToClear, snake.points[0])
		snake.points = snake.points[1:]
	} else {
		score++
		displayGameScore()
	}
	if isSnakeEatingItself() {
		isGameOver = true
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

// isEggInsideSnake is used to check whether an agg is eaten by the snake.
// If the egg is eaten, the egg coordinates equal to one of snake body coordinates.
// Only the egg is inside a snake, the snake will grow up.
func isEggInsideSnake() bool {
	for _, snakeCoordinate := range snake.points {
		if snakeCoordinate.x == egg.point.x && snakeCoordinate.y == egg.point.y {
			return true
		}
	}
	return false
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

func updateEgg() {
	for isEggInsideSnake() {
		coordinatesToClear = append(coordinatesToClear, egg.point)
		egg.point.x, egg.point.y = getNewEggCoordinate()
	}
}

func updateSpeed() {
	_ = getSpeedLevel()
	displaySpeedLevel()
}

func updateGameState() {
	if isGamePaused {
		return
	}
	clearScreen()
	updateSnake()
	updateSpeed()
	updateEgg()
}

func transformCoordinateInsideFrame(coordinate *Coordinate) {
	frameOriginX, frameOriginY := getFrameTopLeftCoordinate()
	frameOriginX += 1
	frameOriginY += 1
	coordinate.x += frameOriginX
	coordinate.y += frameOriginY
	for coordinate.x >= frameOriginX+FrameWidth {
		coordinate.x--
	}
	for coordinate.y >= frameOriginY+FrameHeight {
		coordinate.y--
	}
}

// initGameObj is used to set initial values for the snake variable and the egg variable
// horizontalDirect and verticalDirect is used to determine the snake init direction.
// horizontalDirect=1, verticalDirect=0 means leftward.
// horizontalDirect=0, verticalDirect=-1 means upward.
// horizontalDirect=0, verticalDirect=1 means downward.
// horizontalDirect=-1, verticalDirect=0 means rightward.
func initGameObj() {
	snake = &Snake{
		points:           getInitialSnakeCoordinates(),
		horizontalDirect: 1,
		verticalDirect:   0,
		symbol:           SnakeSymbol,
	}

	egg = &Egg{
		point:  getInitialEggCoordinates(),
		symbol: EggSymbol,
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

func getInitialEggCoordinates() *Coordinate {
	eggInitialCoordinate := &Coordinate{FrameWidth / 2, FrameHeight / 2}
	transformCoordinateInsideFrame(eggInitialCoordinate)

	return eggInitialCoordinate
}

func initScreen() {
	encoding.Register()
	var err error
	Screen, err = tcell.NewScreen()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = Screen.Init(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Using tcell set different colors in background and foreground
	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	Screen.SetStyle(defStyle)
	screenWidth, screenHeight = Screen.Size()
	fmt.Println("screenWidth and screenHeight are", screenWidth, screenHeight)

	if screenWidth < FrameWidth || screenHeight < FrameHeight {
		fmt.Printf("The game frame is defined with %d width and %d height. Increase terminal size and try again ", FrameWidth, FrameHeight)
		os.Exit(1)
	}
}

func printScreen(x, y, borderThickness int, style tcell.Style, char rune) {
	for i := 0; i < borderThickness; i++ {
		for j := 0; j < borderThickness; j++ {
			Screen.SetContent(x+i, y+j, char, nil, style)
		}
	}
}

func getFrameTopLeftCoordinate() (int, int) {
	return (screenWidth-FrameWidth)/2 - 1, (screenHeight-FrameHeight)/2 - 1
}

// displayFrame is used to draw the game's border.
func displayFrame() {
	printUnfilledRectangle(getFrameTopLeftCoordinate())
	Screen.Show()
}

func displayGameObjects() {
	displaySnake()
	displayEgg()
	Screen.Show()
}

func displaySnake() {
	style := tcell.StyleDefault.Foreground(tcell.ColorDarkGreen.TrueColor())
	for _, snakeCoordinate := range snake.points {
		printScreen(snakeCoordinate.x, snakeCoordinate.y, 1, style, snake.symbol)
	}
}

func displayEgg() {
	style := tcell.StyleDefault.Foreground(tcell.ColorOrange.TrueColor())
	printScreen(egg.point.x, egg.point.y, 1, style, egg.symbol)
}

func displayGamePausedInfo() {
	_, frameY := getFrameTopLeftCoordinate()
	printAtCenter(frameY-2, "Game Paused !!", true)
	printAtCenter(frameY-1, "Press p or Enter to resume", true)
}

func displayGameOverInfo() {
	centerY := (screenHeight - FrameHeight) / 2
	printAtCenter(centerY-1, "Game Over !!", false)
	printAtCenter(centerY, fmt.Sprintf("Your Score : %d", score), false)
}

func displayGameScore() {
	_, frameY := getFrameTopLeftCoordinate()
	printAtCenter(frameY+FrameHeight+2, fmt.Sprintf("Current Score : %d", score), false)
}

func displaySpeedLevel() {
	_, frameY := getFrameTopLeftCoordinate()
	printAtCenter(frameY+FrameHeight+3, fmt.Sprintf("Current Speed : %d", speed), false)
}

// printAtCenter is used to print passed content at center horizontally, while vertical coordinate is passed to function.
func printAtCenter(startY int, content string, trackClear bool) {
	startX := (screenWidth - len(content)) / 2
	for i := 0; i < len(content); i++ {
		printScreen(startX+i, startY, 1, tcell.StyleDefault, rune(content[i]))
		if trackClear {
			coordinatesToClear = append(coordinatesToClear, &Coordinate{startX + i, startY})
		}
	}
	Screen.Show()
}

// clearScreen is used to clear necessary coordinates represented by variable coordinatesToClear
func clearScreen() {
	for _, coordinate := range coordinatesToClear {
		printScreen(coordinate.x, coordinate.y, 1, tcell.StyleDefault, ' ')
	}
}

func printUnfilledRectangle(xOrigin, yOrigin int) {
	var upperBorder, lowerBorder rune
	verticalBorder := FrameBorderVertical
	for i := 0; i < FrameWidth; i++ {
		if i == 0 {
			upperBorder = FrameBorderTopLeft
			lowerBorder = FrameBorderBottomLeft
		} else if i == FrameWidth-1 {
			upperBorder = FrameBorderTopRight
			lowerBorder = FrameBorderBottomRight
		} else {
			upperBorder = FrameBorderHorizontal
			lowerBorder = FrameBorderHorizontal
		}
		printScreen(xOrigin+i, yOrigin, FrameBorderThickness, tcell.StyleDefault, upperBorder)
		printScreen(xOrigin+i, yOrigin+FrameHeight, FrameBorderThickness, tcell.StyleDefault, lowerBorder)
	}

	// side boundary
	for i := 1; i < FrameHeight; i++ {
		printScreen(xOrigin, yOrigin+i, FrameBorderThickness, tcell.StyleDefault, verticalBorder)
		printScreen(xOrigin+FrameWidth-1, yOrigin+i, FrameBorderThickness, tcell.StyleDefault, verticalBorder)
	}
}
