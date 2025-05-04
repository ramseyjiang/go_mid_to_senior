package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var serviceName string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new microservice",
	Run: func(cmd *cobra.Command, args []string) {
		if serviceName == "" {
			cmd.PrintErrln("Error: service name required")
			return
		}
		generateService(serviceName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&serviceName, "name", "n", "", "Name of the service")
}

func generateService(name string) {
	destDir := filepath.Join("services", name)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		panic(err)
	}

	// Walk template directory and copy files
	filepath.Walk("templates/service", func(src string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath := strings.TrimPrefix(src, "templates/service")
		destPath := filepath.Join(destDir, relPath)
		if info.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}
		content, err := os.ReadFile(src)
		if err != nil {
			return err
		}
		// Replace placeholders
		content = []byte(strings.ReplaceAll(string(content), "{{.ServiceName}}", name))
		return os.WriteFile(destPath, content, 0644)
	})
}
