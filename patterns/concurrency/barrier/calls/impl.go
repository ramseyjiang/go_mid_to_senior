package calls

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type HttpClient interface {
	Get(url string) (*http.Response, error)
}

type Response struct {
	Url  string
	Body string
}

func performRequests(w io.Writer, client HttpClient, urls []string) {
	var wg sync.WaitGroup
	responseChannel := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := client.Get(url)
			if err != nil {
				responseChannel <- fmt.Sprintf("Error making GET request to %s: %v", url, err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				responseChannel <- fmt.Sprintf("Error reading body of response from %s: %v", url, err)
				return
			}

			if resp.StatusCode != http.StatusOK {
				responseChannel <- fmt.Sprintf("Non-ok status returned from %s: %d", url, resp.StatusCode)
				return
			}

			responseChannel <- fmt.Sprintf("Response from %s:\n%s", url, string(body))
		}(url)
	}

	go func() {
		wg.Wait()
		close(responseChannel)
	}()

	var combinedResponse string
	for response := range responseChannel {
		combinedResponse += response + "\n"
	}

	fmt.Fprint(w, combinedResponse)
}
