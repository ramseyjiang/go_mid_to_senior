package calls

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Barrier struct {
	total int
	count int
	mutex sync.Mutex
	cond  *sync.Cond
}

func NewBarrier(count int) *Barrier {
	b := &Barrier{
		total: count,
		count: 0,
	}
	b.cond = sync.NewCond(&b.mutex)
	return b
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count++
	if b.count >= b.total {
		b.count = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.mutex.Unlock()
}

func makeRequest(url string, barrier *Barrier, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error making request to %s: %v", url, err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading response body: %v", err)
		return
	}
	ch <- string(body)

	barrier.Wait()
}
