package suite

import (
	"fmt"
)

type Stats struct {
	Name      string
	Minutes   float32
	Points    int8
	Rebounds  int8
	Assists   int8
	Turnovers int8
}

func hadAGoodGame(stats Stats) (bool, error) {
	if stats.Assists < 0 || stats.Points < 0 || stats.Rebounds < 0 || stats.Minutes < 0 || stats.Turnovers < 0 {
		return false, fmt.Errorf("stat lines cannot be negative")
	}
	if stats.Name == "" {
		return false, fmt.Errorf("missing player name")
	}
	if stats.Assists >= (stats.Turnovers * 2) {
		return true, nil
	}
	if stats.Assists >= 10 && stats.Rebounds >= 10 && stats.Points >= 10 {
		return true, nil
	} else if stats.Points < 10 && stats.Assists < 10 && stats.Minutes > 25.0 {
		return false, nil
	}
	return false, nil
}
