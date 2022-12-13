package snake

import (
	"fmt"
	"time"
)

func GetSpeedLevel() time.Duration {
	speed1 := 120 * time.Millisecond
	speed2 := 90 * time.Millisecond
	speed3 := 60 * time.Millisecond
	speed4 := 30 * time.Millisecond
	scoreLevel1 := 1
	scoreLevel2 := 2
	scoreLevel3 := 3

	if game.Score > scoreLevel1 && game.Score <= scoreLevel2 {
		game.Speed = 1
		return speed2
	}

	if game.Score > scoreLevel2 && game.Score <= scoreLevel3 {
		game.Speed = 2
		return speed3
	}

	if game.Score > scoreLevel3 {
		game.Speed = 3
		return speed4
	}

	return speed1
}

func updateSpeed() {
	_ = GetSpeedLevel()
	DisplaySpeedLevel()
}

func DisplaySpeedLevel() {
	_, frameY := getFrameTopLeftCoordinate()
	showNoticeScreenCenter(frameY+FrameHeight+3, fmt.Sprintf("Current Speed : %d", game.Speed), false)
}
