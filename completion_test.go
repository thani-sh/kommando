package kommando_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCompletion(t *testing.T) {
	// Binary is built and copied by Makefile build target before tests run
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	exampleDir := filepath.Join(cwd, "example")
	binaryPath := filepath.Join(exampleDir, "kommando")

	tests := []struct {
		name     string
		args     []string
		expected []string
	}{
		{
			name:     "root completion empty",
			args:     []string{"--kommando-completion", ""},
			expected: []string{"add", "kommando", "say", "sub"},
		},
		{
			name:     "root completion partial 's'",
			args:     []string{"--kommando-completion", "s"},
			expected: []string{"say", "sub"},
		},
		{
			name:     "root completion partial 'a'",
			args:     []string{"--kommando-completion", "a"},
			expected: []string{"add"},
		},
		{
			name:     "nested completion 'say '",
			args:     []string{"--kommando-completion", "say", ""},
			expected: []string{"hello"},
		},
		{
			name:     "nested completion 'say h'",
			args:     []string{"--kommando-completion", "say", "h"},
			expected: []string{"hello"},
		},
		{
			name:     "nested completion no match",
			args:     []string{"--kommando-completion", "say", "x"},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binaryPath, tt.args...)
			cmd.Dir = exampleDir
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Command failed: %v\nOutput: %s", err, output)
			}

			// Split output by newline and filter empty lines
			lines := strings.Split(strings.TrimSpace(string(output)), "\n")
			var got []string
			for _, line := range lines {
				trimmed := strings.TrimSpace(line)
				if trimmed != "" {
					got = append(got, trimmed)
				}
			}

			if len(got) != len(tt.expected) {
				t.Errorf("Expected %d matches, got %d. Expected: %v, Got: %v", len(tt.expected), len(got), tt.expected, got)
				return
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Match mismatch at index %d. Expected: %q, Got: %q", i, tt.expected[i], got[i])
				}
			}
		})
	}
}

func TestInitScripts(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	exampleDir := filepath.Join(cwd, "example")
	binaryPath := filepath.Join(exampleDir, "kommando")

	tests := []struct {
		name          string
		args          []string
		expectedParts []string
	}{
		{
			name: "bash init",
			args: []string{"--kommando-init-bash"},
			expectedParts: []string{
				"_kommando_completion()",
				"complete -F _kommando_completion kommando",
				binaryPath, // Should contain the full path
			},
		},
		{
			name: "zsh init",
			args: []string{"--kommando-init-zsh"},
			expectedParts: []string{
				"#compdef kommando",
				"_kommando()",
				"compdef _kommando kommando",
				binaryPath, // Should contain the full path
			},
		},
		{
			name: "bash init with alias",
			args: []string{"--kommando-init-bash", "myalias"},
			expectedParts: []string{
				"_myalias_completion()",
				"complete -F _myalias_completion myalias",
				binaryPath,
			},
		},
		{
			name: "zsh init with alias",
			args: []string{"--kommando-init-zsh", "myalias"},
			expectedParts: []string{
				"#compdef myalias",
				"_myalias()",
				"compdef _myalias myalias",
				binaryPath,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binaryPath, tt.args...)
			cmd.Dir = exampleDir
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Command failed: %v\nOutput: %s", err, output)
			}

			outStr := string(output)
			for _, part := range tt.expectedParts {
				if !strings.Contains(outStr, part) {
					t.Errorf("Output missing expected part %q. Output:\n%s", part, outStr)
				}
			}
		})
	}
}
