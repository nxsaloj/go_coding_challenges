# Challenge 1 | Write your own wc tool

First challenge of the challenge series by John Cricket https://codingchallenges.fyi/challenges/challenge-wc

## Description

This WC tool is written using golang under ccwc. This tool is used to count the number of bytes, lines, words and multibytes in a given file and from the standard input.

## Usage
Use `go run main.go` to run the program

The following options are supported:

- -c: prints the number of bytes in the given file or stdin
- -w: prints the number of words in the given file or stdin
- -l: prints the number of lines in the given file or stdin
- -m: prints the number of multibytes in the given file or stdin

The tool can also be used in stdin mode as follows:

- `printf "content" | go run main.go [option]`
- `cat filename | go run main.go [option]`

Run tests
To run the tests go to the root ccwc repository and run the following command:

`go test ./..`

All unit test under /cmd and integration tests under /test will be executed.