package calls

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

type MockHttpSuccessClient struct{}

func (m *MockHttpSuccessClient) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte(url))),
	}, nil
}

type MockHttpErrorClient struct{}

func (m *MockHttpErrorClient) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: 404,
		Body:       io.NopCloser(bytes.NewBufferString("")),
	}, nil
}

type MockHttpTimeoutClient struct{}

func (m *MockHttpTimeoutClient) Get(url string) (*http.Response, error) {
	// Simulate a delay
	select {
	case <-time.After(3 * time.Second):
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader([]byte(url))),
		}, nil
	}
}

func TestPerformRequests(t *testing.T) {
	t.Run("Test with valid URLs", func(t *testing.T) {
		client := &MockHttpSuccessClient{}
		urls := []string{"http://example.com", "http://example.org"}

		var buf bytes.Buffer
		performRequests(&buf, client, urls)

		got := buf.String()

		if !strings.Contains(got, "Response from http://example.com:\nhttp://example.com") ||
			!strings.Contains(got, "Response from http://example.org:\nhttp://example.org") {
			t.Errorf("performRequests() returned unexpected output: %q", got)
		}
	})

	t.Run("Test with wrong URL", func(t *testing.T) {
		client := &MockHttpErrorClient{}
		urls := []string{"http://wrong.url"}

		var buf bytes.Buffer
		performRequests(&buf, client, urls)

		got := buf.String()

		if !strings.Contains(got, "404") {
			t.Errorf("performRequests() did not handle error correctly, got: %q", got)
		}
	})

	t.Run("Test with timeout", func(t *testing.T) {
		client := &MockHttpTimeoutClient{}
		urls := []string{"http://timeout.url"}

		var buf bytes.Buffer
		start := time.Now()
		performRequests(&buf, client, urls)
		elapsed := time.Since(start)

		if elapsed < 3*time.Second {
			t.Errorf("performRequests() didn't timeout as expected, elapsed: %v", elapsed)
		}
	})
}
