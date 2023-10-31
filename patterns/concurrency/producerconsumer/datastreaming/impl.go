package datastreaming

import "sync"

type Item struct {
	Data int
}

type Streamer struct {
	ch chan Item
	wg *sync.WaitGroup
}

func NewStreamer(bufferSize int) *Streamer {
	return &Streamer{
		ch: make(chan Item, bufferSize),
		wg: &sync.WaitGroup{},
	}
}

func (s *Streamer) Producer(items []Item) {
	s.wg.Add(1)
	go func() {
		for _, item := range items {
			s.ch <- item
		}
		s.wg.Done()
		close(s.ch) // Close the channel after all items are sent.
	}()
}

func (s *Streamer) Consumer(processedItems *[]Item) {
	s.wg.Add(1)
	go func() {
		for item := range s.ch {
			// Simulate processing
			item.Data *= 2
			*processedItems = append(*processedItems, item)
		}
		s.wg.Done()
	}()
}

func (s *Streamer) Close() {
	s.wg.Wait()
}
