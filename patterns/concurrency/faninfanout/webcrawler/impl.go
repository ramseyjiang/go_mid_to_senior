package webcrawler

import (
	"io"
	"net/http"
	"sync"
)

// Fetcher defines an interface for fetching URLs.
type Fetcher interface {
	Fetch(url string) (string, error)
}

// HTTPFetcher implements the Fetcher interface using the http package.
type HTTPFetcher struct{}

// Fetch is the worker function.
// It makes an HTTP GET request to the specified URL and returns the response body as a string.
func (f *HTTPFetcher) Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Crawl uses the fan-out fan-in pattern to fetch data from multiple URLs.
func Crawl(urls []string, fetcher Fetcher) ([]string, []error) {
	var wg sync.WaitGroup
	results := make(chan string, len(urls))
	errs := make(chan error, len(urls))

	// Fan-out
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			body, err := fetcher.Fetch(url)
			if err != nil {
				errs <- err
				return
			}
			results <- body
		}(url)
	}

	// Fan-in
	go func() {
		wg.Wait()
		close(results)
		close(errs)
	}()

	var bodies []string
	var errors []error
	for {
		select {
		case body, ok := <-results:
			if ok {
				bodies = append(bodies, body)
			} else {
				results = nil
			}
		case err, ok := <-errs:
			if ok {
				errors = append(errors, err)
			} else {
				errs = nil
			}
		}
		if results == nil && errs == nil {
			break
		}
	}

	return bodies, errors
}
