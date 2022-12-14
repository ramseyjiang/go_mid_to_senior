package mergechans

import (
	"sync"
)

func joinChannels(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(chs))
		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for num := range ch {
					mergedCh <- num
				}
			}(ch, wg)
		}
		wg.Wait()
		close(mergedCh)
	}()
	return mergedCh
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
