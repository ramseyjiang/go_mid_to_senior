package cmd

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check health of all services",
	Run: func(cmd *cobra.Command, args []string) {
		checkHealth()
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}

func checkHealth() {
	services := []string{"service1:8081", "service2:8082", "service3:8083"} // Example
	var wg sync.WaitGroup
	results := make(chan string, len(services))

	for _, addr := range services {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			resp, err := http.Get(fmt.Sprintf("http://%s/health", addr))
			if err != nil {
				results <- fmt.Sprintf("%s: DOWN (%v)", addr, err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				results <- fmt.Sprintf("%s: UP", addr)
			} else {
				results <- fmt.Sprintf("%s: DOWN (status %d)", addr, resp.StatusCode)
			}
		}(addr)
	}

	wg.Wait()
	close(results)
	for result := range results {
		fmt.Println(result)
	}
}
