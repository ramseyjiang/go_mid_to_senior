package main

import "fmt"

// original code lots of mistakes.

// parameter id int is not used, so I delete it in my code.
func workers(id int, jobs <-chan int, results chan<- int) {
	// here should not use for range. It will make channel block, because it will wait the info in channel jobs forever.
	// I have changed it to for select.
	// In my code, if it has jobs come, it will run. If not, it will use the default, it won't block anymore
	for j := range jobs {
		// here j does not pass into the goroutine, it will make the j value always the same one, it is not what we want.
		// Hence, I add j int into the go func() in my code.
		// In my code, I also add default case in the switch, if not, it will channel close first sometimes.
		go func() {
			switch j % 3 {
			case 0:
				j = j * 1
			case 1:
				j = j * 2
				results <- j * 2
			case 2:
				results <- j * 3
				j = j * 3
			}
		}()
	}
}

func main() {
	// As it will output a sum, channels should have buffer, and the buffer number should be the same with total num, I explain it below.
	jobs := make(chan int)
	results := make(chan int)

	// here 1000000000 is a billion, it's huge for goroutine.
	// In golang, each goroutine size minimum is 2k, so it means it will use at least 2kb memory.
	// If a billion goroutine, it will use 2000G memory, in my mac pro, the maximum number I can run is 10000000, as 10 million.
	// I define a const named total at the top of my code, "const total = 10000000"
	// After that, I use total to replace everywhere I should use the number.
	// In my code, I also use sync.WaitGroup to make sure the code won't have race condition and deadlock.
	for i := 1; i <= 1000000000; i++ {
		// here is the same mistake as the first one, so I add i int into go func() goroutine.
		go func() {
			if i%2 == 0 {
				i += 99
			}
			jobs <- i
		}()
	}
	close(jobs)

	// I have removed the jobs2 logic.
	// Because I think it is useless. It is just a slice and have 1000 length limit. why not use 1000 to do a for loop directly?
	jobs2 := []int{}
	for w := 1; w < 1000; w++ {
		jobs2 = append(jobs2, w)
	}

	// I replace for range to for loop, because it is stable length 1000. And for loop directly is high efficiency than for range.
	// This link is my blog on medium, I have explained and tested the for range and for loop one year ago.
	// https://medium.com/@ramseyjiang_22278/in-my-previous-article-i-mentioned-there-are-6-simple-ways-to-help-you-to-optimise-golang-code-81c013f70ca4
	for i, w := range jobs2 {
		go workers(w, jobs, results)
		i = i + 1
	}
	close(results)

	var sum int32 = 0
	// here will block all the time. I use for select to fix it.
	// In the for select, I don't use default return at here.
	// Because results is a channel and it runs concurrent. If I use return, when current all data has been taken, it will return.
	// It won't care the channel is nothing clear or not.
	// Hence, I remove the default return case, it will make all data in the results channel can be taken out.
	for r := range results {
		sum += int32(r)
	}

	// at the end, in my code, I let it sleep 5 seconds to make sure all goroutines have been finished.
	fmt.Println(sum)
}
