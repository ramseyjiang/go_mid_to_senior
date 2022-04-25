package main

import (
	"fmt"
	"log"
)

// sq will return the square of int.
func sq(a int) int {
	return a * a
}

type MyFunc func(x int) int

// Decorator needs to return the same type of function it will alter, so that the new function will work with the original code seamlessly
func main() {
	// addLogger is decorator function for MyFunc
	addLogger := func(fn MyFunc) MyFunc {

		// return this new altered function
		return func(x int) int {
			log.Printf("Here is the function of type %T will spit out some result: ", fn)

			// Since this example uses a function signature with return value, it needs to return the call of the original function
			return fn(x)
			// If we were implementing a middleware for http handlers from standard library,
			// then we would simply call a handler function without returning, since it has no return value.
		}
	}

	// construct a new function instance, or in other words, decorate or apply a middleware function to sq function.
	logAddedSq := addLogger(sq)

	logAddedSq(5) // if it does not use fmt.Println(), it won't output any results.
	fmt.Println(logAddedSq(5))
}
