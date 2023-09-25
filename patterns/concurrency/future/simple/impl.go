package simple

import (
	"fmt"
	"time"
)

type Future struct {
	result chan float64
	err    chan error
}

func CalculateAverage(data []int) *Future {
	f := &Future{
		result: make(chan float64),
		err:    make(chan error),
	}

	go func() {
		if len(data) == 0 {
			f.err <- fmt.Errorf("data slice is empty")
			return
		}

		var sum int
		for _, v := range data {
			sum += v
		}
		avg := float64(sum) / float64(len(data))

		// Simulating some processing time
		time.Sleep(2 * time.Second)

		f.result <- avg
	}()

	return f
}

func (f *Future) Get() (float64, error) {
	select {
	case res := <-f.result:
		return res, nil
	case err := <-f.err:
		return 0, err
	}
}
