package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet a user",
	Long:  `Send greetings to the specified user`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	// Add command line param
	rootCmd.AddCommand(greetCmd)
	greetCmd.Flags().StringP("name", "n", "Anonymous", "Specify user name")
}
