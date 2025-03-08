package cmd

import "testing"

// strToBytes helper function to convert string input to []byte
func strToBytes(value string) []byte {
	return []byte(value)
}

func TestCount(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		commands []string
		expected []int
	}{
		{
			name:     "Count bytes",
			content:  "hello",
			commands: []string{"c"},
			expected: []int{5},
		},
		{
			name:     "Count lines",
			content:  "line1\nline2\nline3\n",
			commands: []string{"l"},
			expected: []int{3},
		},
		{
			name:     "Count words",
			content:  "hello world, this is Go",
			commands: []string{"w"},
			expected: []int{5},
		},
		{
			name:     "Count multibytes",
			content:  "héllo", // é is a multibyte character
			commands: []string{"m"},
			expected: []int{5},
		},
		{
			name:     "Count multiple metrics",
			content:  "Go is awesome\nLet's code\n",
			commands: []string{"l", "w", "c"},
			expected: []int{2, 5, 25},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := count(strToBytes(test.content), test.commands)
			if err != nil {
				t.Fatalf("count() returned an error: %v", err)
			}

			if len(result) != len(test.expected) {
				t.Errorf("%s: expected %d, got %d", test.name, len(test.expected), len(result))
			}
			for i, v := range test.expected {
				if result[i] != v {
					t.Errorf("%s: expected %d for command %s, got %d", test.name, v, test.commands[i], result[i])
				}
			}
		})
	}
}
