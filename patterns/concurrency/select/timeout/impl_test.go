package timeout

import (
	"fmt"
	"testing"
	"time"
)

func TestFetchDataWithTimeout(t *testing.T) {
	resultChan := make(chan Result)

	t.Run("returns data when the fetch finishes before the timeout", func(t *testing.T) {
		go fetchDataWithTimeout(resultChan, 1*time.Second, 2*time.Second)

		result := <-resultChan
		if result.Error != nil {
			t.Errorf("Expected no error, got %v", result.Error)
		}

		if result.Data != "Data fetched" {
			t.Errorf("Expected 'Data fetched', got %s", result.Data)
		}
		fmt.Println(result.Data)
	})

	t.Run("times out when the fetch does not finish before the timeout", func(t *testing.T) {
		go fetchDataWithTimeout(resultChan, 2*time.Second, 1*time.Second)

		result := <-resultChan
		if result.Error == nil {
			t.Error("Expected a timeout error, but got no error")
		}

		if result.Error.Error() != "fetch timed out" {
			t.Errorf("Expected 'fetch timed out', got %v", result.Error)
		}
		fmt.Println(result.Error.Error())
	})
}
