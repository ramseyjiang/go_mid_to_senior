package calls

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("Wait", func(t *testing.T) {
		b := NewBarrier(total)
		var wg sync.WaitGroup

		for i := 0; i < total; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b.Wait()
			}()
		}

		wg.Wait()
		if b.count != 0 {
			t.Errorf("Expected count to be 0, got %d", b.count)
		}
	})
}

func TestMakeRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	t.Run("Successful request", func(t *testing.T) {
		b := NewBarrier(1)
		var wg sync.WaitGroup
		ch := make(chan string, 1)

		wg.Add(1)
		go makeRequest(ts.URL, b, &wg, ch)

		wg.Wait()
		close(ch)

		msg := <-ch
		if msg != "Hello, client\n" {
			t.Errorf("Expected 'Hello, client\\n', got '%s'", msg)
		}
	})

	t.Run("Failed request", func(t *testing.T) {
		b := NewBarrier(1)
		var wg sync.WaitGroup
		ch := make(chan string, 1)

		wg.Add(1)
		go makeRequest("http://invalid.url", b, &wg, ch)

		wg.Wait()
		close(ch)

		msg := <-ch
		if msg == "" {
			t.Error("Expected an error message, got an empty string")
		}
	})
}
