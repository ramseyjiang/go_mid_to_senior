package twotestways

import (
	"fmt"
)

type doughnutsBox struct {
	capacity  int
	doughnuts []string
}

var knownDoughnutTypes = map[string]bool{
	"Mocha Tea":                 true,
	"Lime & Coconut (ve)":       true,
	"Home Made Raspberry Jam":   true,
	"Cinnamon Scroll (ve)":      true,
	"Sri Lankan Cinnamon Sugar": true,
}

func newDoughnutsBox(capacity int) *doughnutsBox {
	return &doughnutsBox{
		capacity:  capacity,
		doughnuts: make([]string, 0),
	}
}

func (b *doughnutsBox) pack(doughnuts []string) (numOfDoughnutsInTheBox int, err error) {
	unrecognizedItems := make([]string, 0)

	if len(doughnuts) > b.capacity {
		return 0, fmt.Errorf("failed to put %d doughnuts in the box, it's only has %d doughnuts capacity", len(doughnuts), b.capacity)
	}

	for _, doughnut := range doughnuts {
		if _, found := knownDoughnutTypes[doughnut]; found {
			b.doughnuts = append(b.doughnuts, doughnut)
			continue
		}
		unrecognizedItems = append(unrecognizedItems, doughnut)
	}

	if len(unrecognizedItems) > 0 {
		err = fmt.Errorf("the following items cannot be placed into the box: %v", unrecognizedItems)
	}
	return len(b.doughnuts), err
}
