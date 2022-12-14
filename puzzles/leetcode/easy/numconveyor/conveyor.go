package numconveyor

func conveyor() (results []int) {
	receiveNums := make(chan int)
	squareResults := make(chan int)

	go func() {
		for x := 0; x <= 4; x++ {
			receiveNums <- x
		}
		close(receiveNums)
	}()

	go func() {
		for x := range receiveNums {
			squareResults <- x * x
		}
		close(squareResults)
	}()

	for res := range squareResults {
		results = append(results, res)
	}

	return
}

// contains is used to check an integer element is in a slice or not.
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
