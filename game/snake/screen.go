package snake

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

// screen is used to hold screen information that comes from tcell package
var screen tcell.Screen

// screenWidth and screenHeight are used to hold screen's full width and height respectively.
var screenWidth, screenHeight int

// InitGameObj is used to set initial values for the snake variable and the egg variable
// xDirect and yDirect is used to determine the snake init direction.
// xDirect=1, yDirect=0 means leftward.
// xDirect=0, yDirect=-1 means upward.
// xDirect=0, yDirect=1 means downward.
// xDirect=-1, yDirect=0 means rightward.
func InitGameObj() {
	InitScreen()
	snake = &Snake{
		points:  getInitialSnakeCoordinates(),
		xDirect: 1,
		yDirect: 0,
		symbol:  SymbolSnake,
	}

	egg = &Egg{
		pos:    getInitialEggCoordinates(),
		symbol: SymbolEgg,
	}

	DisplayFrame()
	DisplayGameScore()
	DisplaySpeedLevel()
}

func InitScreen() {
	screen, err = tcell.NewScreen()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = screen.Init(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Using tcell set different colors in background and foreground
	defaultStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(defaultStyle)
	screenWidth, screenHeight = screen.Size()

	if screenWidth < FrameWidth || screenHeight < FrameHeight {
		fmt.Printf("The game frame is defined with %d width and %d height. Increase terminal size and try again ", FrameWidth, FrameHeight)
		os.Exit(1)
	}
}

// DisplayFrame is used to draw the game's border.
func DisplayFrame() {
	drawPlayArea(getFrameTopLeftCoordinate())
	screen.Show()
}

func drawPlayArea(xOrigin, yOrigin int) {
	var upperBorder, lowerBorder int32
	verticalBorder := FrameBorderVertical
	for i := 0; i < FrameWidth; i++ {
		switch i {
		case 0:
			upperBorder = FrameBorderHorizontal
			lowerBorder = FrameBorderVertical
		case FrameWidth - 1:
			upperBorder = FrameBorderHorizontal
			lowerBorder = FrameBorderVertical
		default:
			upperBorder = FrameBorderHorizontal
			lowerBorder = FrameBorderHorizontal
		}
		drawElement(xOrigin+i, yOrigin, FrameBorderThickness, tcell.StyleDefault, upperBorder)               // print top border
		drawElement(xOrigin+i, yOrigin+FrameHeight-1, FrameBorderThickness, tcell.StyleDefault, lowerBorder) // print bottom border
	}

	// side boundary
	for i := 1; i < FrameHeight; i++ {
		drawElement(xOrigin, yOrigin+i, FrameBorderThickness, tcell.StyleDefault, verticalBorder)              // print left side border
		drawElement(xOrigin+FrameWidth-1, yOrigin+i, FrameBorderThickness, tcell.StyleDefault, verticalBorder) // print right side border
	}
}

func DisplayGameScore() {
	_, frameY := getFrameTopLeftCoordinate()
	showNoticeScreenCenter(frameY+FrameHeight+2, fmt.Sprintf("Current Score : %d", score), false)
}

func DisplayGamePausedInfo() {
	_, frameY := getFrameTopLeftCoordinate()
	showNoticeScreenCenter(frameY+4, "Game Paused ! Press r to restart.", true)
	showNoticeScreenCenter(frameY+3, "Press p or Enter to resume", true)
}

func DisplayGameOverInfo() {
	centerY := (screenHeight - FrameHeight) / 2
	showNoticeScreenCenter(centerY-1, "Game Over !!", false)
	showNoticeScreenCenter(centerY, fmt.Sprintf("Your Score : %d", score), false)
}
