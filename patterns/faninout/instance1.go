package faninout

import (
	"fmt"
)

func EntryInstance1() {
	randomNumbers := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	// generate the common channel with inputs
	inputChan := generateInputChan(randomNumbers)

	// Fan-out to 2 Goroutine
	c1 := squareNumber(inputChan)
	c2 := squareNumber(inputChan)

	// Fan-in the resulting squared numbers
	c := fanIn(c1, c2)
	sum := 0

	// Do the summation
	for i := 0; i < len(randomNumbers); i++ {
		sum += <-c
	}
	fmt.Println("total Sum of Squares: ", sum)
}

func generateInputChan(numbers []int) <-chan int {
	inputChan := make(chan int)
	go func() {
		for _, n := range numbers {
			inputChan <- n
		}
		close(inputChan)
	}()
	return inputChan
}

// Each instance of the squareNumber function reads from the same input channel until that channel is closed.
func squareNumber(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// The fanIn function can read from multiple inputs channels and proceed until all are closed by multiplexing the input channels
// onto a single channel that’s closed when all the inputs are closed.
func fanIn(input1, input2 <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
