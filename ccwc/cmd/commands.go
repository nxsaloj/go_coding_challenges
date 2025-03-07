package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute runs cobra default defined command.
// If an error occurs, it prints the error message and exits the program.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// rootCmd represents the root command for the application.
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "Ccwc is a word count tool",
	Long: `A word count tool based on unix wc tool built
for learning Golang by nxcrypt.`}
