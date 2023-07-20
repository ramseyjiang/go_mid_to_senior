package simple

import "time"

func compute(input int) <-chan int {
	result := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		result <- input * 2
	}()
	return result
}
