package fetch

import (
	"testing"
)

func TestFetch(t *testing.T) {
	testCases := []struct {
		name string
		url  string
	}{
		{name: "Example.com", url: "http://example.com"},
		{name: "Example.org", url: "http://example.org"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := NewFuture(func() (Result, error) {
				return fetch(tc.url)
			})

			result := f.Get()
			if result.Err != nil {
				t.Errorf("Error: %v", result.Err)
			} else {
				// fmt.Println("Result:", string(result.Body))
			}
		})
	}
}
