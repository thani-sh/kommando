package kommando_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Build binary
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get current directory: %v\n", err)
		os.Exit(1)
	}

	exampleDir := filepath.Join(cwd, "example")
	binaryPath := filepath.Join(exampleDir, "kommando")

	// Build the kommando binary
	cmd := exec.Command("go", "build", "-o", binaryPath, "./cmd/kommando/...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to build kommando binary: %v\n", err)
		os.Exit(1)
	}

	// Run tests
	code := m.Run()

	// Cleanup
	os.Remove(binaryPath)

	os.Exit(code)
}

func TestKommandoCLI(t *testing.T) {
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
