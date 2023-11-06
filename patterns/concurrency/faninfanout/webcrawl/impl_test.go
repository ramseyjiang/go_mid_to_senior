package webcrawl

import (
	"testing"
)

// MockFetcher is a mock fetcher for testing.
type MockFetcher struct{}

// Fetch returns a mocked response for a given URL.
func (f *MockFetcher) Fetch(url string) (string, error) {
	return "mocked body of " + url, nil
}

func TestCrawl(t *testing.T) {
	mockFetcher := &MockFetcher{}

	urls := []string{
		"http://example.com",
		"http://example.org",
		"http://example.net",
	}

	t.Run("TestCrawlFunctionality", func(t *testing.T) {
		results, errors := Crawl(urls, mockFetcher)
		if len(errors) > 0 {
			t.Errorf("Crawl() returned errors: %v", errors)
		}
		if len(results) != len(urls) {
			t.Errorf("Crawl() got = %d, want %d", len(results), len(urls))
		}

		expectedResults := make(map[string]bool)
		for _, url := range urls {
			expectedResults["mocked body of "+url] = false
		}

		for _, result := range results {
			if _, ok := expectedResults[result]; ok {
				expectedResults[result] = true
			} else {
				t.Errorf("Crawl() got an unexpected result: %s", result)
			}
		}

		for url, found := range expectedResults {
			if !found {
				t.Errorf("Crawl() did not get a result for: %s", url)
			}
		}
	})
}
