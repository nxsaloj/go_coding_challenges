package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	command := args[0]
	filename := args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("file %v not found\n", filename)
	}
	defer file.Close()

	count, err := count(file, command[1:])
	if err != nil {
		fmt.Printf("An error has occurred %v", err)
	}

	fmt.Printf("%d %v", count, filename)

}

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
