package generator

import "fmt"

// Generator func which produces data which might be computationally expensive.
func fib(n int) chan int {
	c := make(chan int)
	go func() {
		for i, j := 0, 1; i < n; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()
	return c
}

func Entry() {
	// fib returns the fibonacci numbers lesser than 100
	for i := range fib(100) {
		// Consumer which consumes the data produced by the generator, which further does some extra computations
		v := i * i
		fmt.Println(v)
	}
}
