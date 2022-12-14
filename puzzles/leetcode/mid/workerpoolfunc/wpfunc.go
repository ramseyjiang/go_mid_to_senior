package workerpoolfunc

const numJobs = 6

// worker using "f func(int)" as param, the "int" after the "f func(int)" is the return type of "f func(int)"
func worker(id int, f func(int) int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		// fire the response into the resulting channel. f(j) is the multiplier response.
		results <- f(j)
	}
}

func reuseRoutine() (res []int) {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	multiplier := func(x int) int {
		return x * 10
	}

	// for each worker creates a goroutine and wait for a new job, apply the given function "multiplier" to it
	for w := 1; w <= 3; w++ {
		// while not creating a new goroutine every time, but simply reusing the existing one.
		go worker(w, multiplier, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		res = append(res, <-results)
	}
	close(results)

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
