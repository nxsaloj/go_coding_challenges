package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

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

		var flagsUsed bool = false
		if cmd.Flags().Changed("bytes") {
			flagsUsed = true
			handleCommand(filename, "c")
		}

		if cmd.Flags().Changed("lines") {
			flagsUsed = true
			handleCommand(filename, "l")
		}

		if cmd.Flags().Changed("words") {
			flagsUsed = true
			handleCommand(filename, "w")
		}

		if cmd.Flags().Changed("multibytes") {
			flagsUsed = true
			handleCommand(filename, "m")
		}

		if !flagsUsed {
			commands := []string{"l", "w", "c"}
			handleCommands(filename, commands)
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

// handleFile opens the given file, it uses the count function to return the count.
// If an error occurs, an error message is printed.
//
// Parameters:
//   - filename: The path of the file.
//   - commands: The types of count operations to perform ("c", "l", "w", or "m").
//
// Returns:
//   - An integer array representing the count of the specified unit types.
func handleFile(filename string, commands []string) []int {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:\n", err)
		os.Exit(2)
	}

	counts, err := count(content, commands)
	if err != nil {
		fmt.Printf("An error has occurred %v", err)
	}

	return counts
}

// handleCommand processes the given file with the specified command and prints the result.
//
// Parameters:
//   - filename: The path of the file.
//   - command: The type of count operation to perform ("c", "l", "w", or "m").
func handleCommand(filename string, command string) {
	var count int = 0
	if len(filename) > 0 {
		count = handleFile(filename, []string{command})[0]
	} else {
		count = handleStdin([]string{command})[0]
	}
	fmt.Printf("%d %s", count, filename)
}

// handleStdin reads the standard input, it uses the count function to return the count.
// If an error occurs, an error message is printed.
//
// Parameters:
//   - commands: The types of count operations to perform ("c", "l", "w", or "m").
//
// Returns:
//   - An integer array representing the count of the specified unit types.
func handleStdin(commands []string) []int {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error:\n", err)
		os.Exit(2)
	}

	counts, err := count(content, commands)
	if err != nil {
		fmt.Printf("An error has occurred %v", err)
	}

	return counts
}

// handleCommands processes the given file with the specified commands and prints the results.
//
// Parameters:
//   - filename: The path of the file.
//   - commands: The types of count operation to perform ("c", "l", "w", or "m").
func handleCommands(filename string, commands []string) {
	var counts []int
	if len(filename) > 0 {
		counts = handleFile(filename, commands)
	} else {
		counts = handleStdin(commands)
	}

	countsStr := strings.Trim(fmt.Sprint(counts), "[]")
	fmt.Printf("%s %s", countsStr, filename)
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
func count(content []byte, commands []string) ([]int, error) {
	var counters []int

	for _, command := range commands {
		switch command {
		case "c":
			counters = append(counters, len(content))
		case "l":
			var counter int
			for i := 0; i < len(content); i++ {
				if string(content[i]) == "\n" {
					counter++
				}
			}
			counters = append(counters, counter)
		case "w":
			counters = append(counters, len(bytes.Fields(content)))
		case "m":
			counters = append(counters, len(bytes.Runes(content)))
		}
	}

	return counters, nil
}
