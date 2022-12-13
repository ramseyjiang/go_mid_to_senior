package snake

import "github.com/gdamore/tcell/v2"

func drawElement(x, y, borderThickness int, style tcell.Style, char int32) {
	for i := 0; i < borderThickness; i++ {
		for j := 0; j < borderThickness; j++ {
			screen.SetContent(x+i, y+j, char, nil, style)
		}
	}
}

func getFrameTopLeftCoordinate() (int, int) {
	return (screenWidth-FrameWidth)/2 - 1, (screenHeight-FrameHeight)/2 - 1
}

// showNoticeScreenCenter is used to print passed content at center horizontally, while vertical coordinate is passed to function.
func showNoticeScreenCenter(startY int, content string, trackClear bool) {
	startX := (screenWidth - len(content)) / 2
	for i := 0; i < len(content); i++ {
		drawElement(startX+i, startY, 1, tcell.StyleDefault, int32(content[i]))
		if trackClear {
			coordinatesToClear = append(coordinatesToClear, &Coordinate{startX + i, startY})
		}
	}
	screen.Show()
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
