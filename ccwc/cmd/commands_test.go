package cmd

import (
	"testing"
)

// strToBytes helper function to convert string input to []byte
func strToBytes(value string) []byte {
	return []byte(value)
}

var contents = []string{
	"hello",
	"line1\nline2\nline3\n",
	"hello world, this is Go",
	"héllo",
	"Go is awesome\nLet's code\n",
}

var singleCommandTests = []struct {
	name, command string
	expected      []int
}{
	{
		name:    "Count bytes",
		command: "c",
		expected: []int{
			5,
			18,
			23,
			6,
			25,
		},
	},
	{
		name:    "Count lines",
		command: "l",
		expected: []int{
			0,
			3,
			0,
			0,
			2,
		},
	},
	{
		name:    "Count words",
		command: "w",
		expected: []int{
			1,
			3,
			5,
			1,
			5,
		},
	},
	{
		name:    "Count multibytes",
		command: "m",
		expected: []int{
			5,
			18,
			23,
			5,
			25,
		},
	},
}

var multipleCommandsTest = []struct {
	name     string
	commands []string
	expected [][]int
}{
	{
		name:     "Count multiple metrics",
		commands: []string{"l", "w", "c"},
		expected: [][]int{
			{0, 1, 5},
			{3, 3, 18},
			{0, 5, 23},
			{0, 1, 6},
			{2, 5, 25},
		},
	},
}

func TestCount(t *testing.T) {
	for _, test := range singleCommandTests {
		t.Run(test.name, func(t *testing.T) {
			for index, content := range contents {
				results, err := count(strToBytes(content), []string{test.command})
				if err != nil {
					t.Fatalf("count() returned an error: %v", err)
				}

				if results[0] != test.expected[index] {
					t.Errorf("%s: expected %d got %d, for command %s", test.name, test.expected[index], results[0], test.command)
				}
			}
		})
	}

	for _, test := range multipleCommandsTest {
		t.Run(test.name, func(t *testing.T) {
			for index, content := range contents {
				results, err := count(strToBytes(content), test.commands)
				if err != nil {
					t.Fatalf("count() returned an error: %v", err)
				}

				if len(results) != len(test.expected[index]) {
					t.Errorf("%s: expected %d, got %d", test.name, len(test.expected[index]), len(results))
				}

				for idx, _result := range results {
					if _result != test.expected[index][idx] {
						t.Errorf("%s: expected %d for command %s, got %d", test.name, test.expected[index][idx], test.commands[idx], _result)
					}
				}
			}
		})
	}
}
