package cmd

import (
	"bufio"
	"fmt"
	"io"
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
for learning Golang by nxcrypt.`,
	Run: func(cmd *cobra.Command, args []string) {
		var filename string
		if len(args) > 0 {
			filename = args[0]
		}

		if cmd.Flags().Changed("bytes") {
			handleCommand(filename, "c")
		}

		if cmd.Flags().Changed("lines") {
			handleCommand(filename, "l")
		}

		if cmd.Flags().Changed("words") {
			handleCommand(filename, "w")
		}

		if cmd.Flags().Changed("multibytes") {
			handleCommand(filename, "m")
		}
	},
}

// init initializes the flags for the root command.
// init() is a special function that is automatically called when the program
// starts and it runs before the main()
func init() {
	rootCmd.Flags().BoolP("bytes", "c", false, "count bytes in content")
	rootCmd.Flags().BoolP("lines", "l", false, "count lines in content")
	rootCmd.Flags().BoolP("words", "w", false, "count words in content")
	rootCmd.Flags().BoolP("multibytes", "m", false, "count multibytes in content")
}

// handleCommand processes the given file with the specified command and prints the result.
//
// Parameters:
//   - filename: The path of the file.
//   - command: The type of count operation to perform ("c", "l", "w", or "m").
func handleCommand(filename string, command string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("file %v not found\n", filename)
	}

	defer file.Close()

	count, err := count(file, command)
	if err != nil {
		fmt.Printf("An error has occurred %v", err)
	}

	fmt.Printf("%d %s", count, filename)
}

// count counts the number of units in the content based on the specified command.
//
// This function counts the number of units defined by the `command` parameter using
// a scan method. If an error occurs while scanning, it returns an error. Otherwise,
// it returns the final count.
//
// Parameters:
//   - content: The content to be scanned, which implements the `io.Reader` interface.
//   - command: The type of count operation to perform ("c", "l", "w", or "m").
//
// Returns:
//   - An integer representing the count of the specified unit type.
//   - An error, if any, encountered during scanning the content.
func count(content io.Reader, command string) (int, error) {
	scanner := bufio.NewScanner(content)

	switch command {
	case "c":
		scanner.Split(bufio.ScanBytes)
	case "l":
		scanner.Split(bufio.ScanLines)
	case "w":
		scanner.Split(bufio.ScanWords)
	case "m":
		scanner.Split(bufio.ScanRunes)
	}

	var counter int
	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("there was an error reading the content %w", err)
	}

	return counter, nil
}
