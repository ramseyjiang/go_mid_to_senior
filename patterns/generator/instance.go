package generator

import "fmt"

// Generator func which produces data which might be computationally expensive.
// Hence, we can do the generation data concurrently.
// That way, the program does not have to wait until all data is generated.

// fib returns a written only channel which transports fibonacci numbers
func fib(len int) <-chan int {
	// make a buffered channel
	c := make(chan int, len)

	// run generation concurrently
	go func() {
		for i, j := 0, 1; i < len; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()

	// return channel
	return c
}

func Entry() {
	// fib returns the fibonacci numbers lesser than 10
	for i := range fib(10) {
		// Consumer which consumes the data produced by the generator, which further does some extra computations
		v := i * i
		fmt.Println("current fib number is ", i)
		fmt.Println(v)
	}
}
