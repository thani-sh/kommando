package kommando_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestKommandoCLI(t *testing.T) {
	// Assume the binary has already been built and copied to example directory
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	
	exampleDir := filepath.Join(cwd, "example")
	binaryPath := filepath.Join(exampleDir, "kommando")

	// Test cases
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "add command",
			args:     []string{"add", "1", "3"},
			expected: "4",
		},
		{
			name:     "sub command",
			args:     []string{"sub", "5", "3"},
			expected: "2",
		},
		{
			name:     "sub command with -i flag",
			args:     []string{"sub", "-i", "5", "3"},
			expected: "-2",
		},
		{
			name:     "say hello command",
			args:     []string{"say", "hello", "Earth"},
			expected: "Hello Earth!",
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

			result := strings.TrimSpace(string(output))
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}
