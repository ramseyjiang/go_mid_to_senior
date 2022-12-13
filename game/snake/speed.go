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

func updateSpeed() {
	_ = GetSpeedLevel()
	DisplaySpeedLevel()
}

func DisplaySpeedLevel() {
	_, frameY := getFrameTopLeftCoordinate()
	showNoticeScreenCenter(frameY+FrameHeight+3, fmt.Sprintf("Current Speed : %d", speed), false)
}
