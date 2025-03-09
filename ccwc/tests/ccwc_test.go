package tests

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestCCWCIntegration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		args     []string
		expected string
	}{
		{
			name:     "Count bytes",
			input:    "hello",
			args:     []string{"-c"},
			expected: "5",
		},
		{
			name:     "Count lines",
			input:    "line1\nline2\nline3\n",
			args:     []string{"-l"},
			expected: "3",
		},
		{
			name:     "Count words",
			input:    "hello world, this is Go",
			args:     []string{"-w"},
			expected: "5",
		},
		{
			name:     "Count multibytes",
			input:    "héllo",
			args:     []string{"-m"},
			expected: "5",
		},
		{
			name:     "Count default empty flag",
			input:    "Go is awesome\nLet's code\n",
			args:     []string{""},
			expected: "2 5 25",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//cmd := exec.Command("go run ../main.go", test.args...)
			cmd := exec.Command("go", append([]string{"run", "../main.go"}, test.args...)...)

			cmd.Stdin = strings.NewReader(test.input)
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				t.Fatalf("command failed: %v", err)
			}

			if strings.TrimSpace(out.String()) != test.expected {
				t.Errorf("%s:\n\texpected %s\n\tgot %s", test.name, test.expected, out.String())
			}

			if strings.Contains(test.args[len(test.args)-1], ".txt") {
				exec.Command("rm", "test.txt").Run()
			}
		})
	}
}

/*
func TestCCWCFileIntegration(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		args     []string
		expected string
	}{
		{
			name:     "Count bytes from file",
			content:  "hello",
			args:     []string{"-c", "testfile.txt"},
			expected: "5 testfile.txt\n",
		},
		{
			name:     "Count lines from file",
			content:  "line1\nline2\nline3\n",
			args:     []string{"-l", "testfile.txt"},
			expected: "3 testfile.txt\n",
		},
		{
			name:     "Count words from file",
			content:  "hello world, this is Go",
			args:     []string{"-w", "testfile.txt"},
			expected: "5 testfile.txt\n",
		},
		{
			name:     "Count multibytes from file",
			content:  "héllo",
			args:     []string{"-m", "testfile.txt"},
			expected: "5 testfile.txt\n",
		},
		{
			name:     "Count multiple metrics from file",
			content:  "Go is awesome\nLet's code\n",
			args:     []string{"-l", "-w", "-c", "testfile.txt"},
			expected: "2 5 22 testfile.txt\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := os.WriteFile("testfile.txt", []byte(test.content), 0644)
			if err != nil {
				t.Fatalf("failed to create test file: %v", err)
			}

			cmd := exec.Command("go run ../cmd/command.go", test.args...)
			var out bytes.Buffer
			cmd.Stdout = &out

			err = cmd.Run()
			if err != nil {
				t.Fatalf("command failed: %v", err)
			}

			if out.String() != test.expected {
				t.Errorf("expected %q, got %q", test.expected, out.String())
			}

			os.Remove("testfile.txt")
		})
	}
}
*/
