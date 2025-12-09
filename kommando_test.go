package kommando_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestKommandoCLI(t *testing.T) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// Build the kommando binary for testing
	buildDir := filepath.Join(cwd, "build")
	binaryBuildPath := filepath.Join(buildDir, "kommando")
	
	// Create build directory if it doesn't exist
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		t.Fatalf("Failed to create build directory: %v", err)
	}

	buildCmd := exec.Command("go", "build", "-o", binaryBuildPath, "./cmd/kommando/...")
	buildCmd.Dir = cwd
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build kommando binary: %v", err)
	}

	// Copy binary to example directory
	exampleDir := filepath.Join(cwd, "example")
	binaryPath := filepath.Join(exampleDir, "kommando")
	
	input, err := os.ReadFile(binaryBuildPath)
	if err != nil {
		t.Fatalf("Failed to read binary: %v", err)
	}
	
	if err := os.WriteFile(binaryPath, input, 0755); err != nil {
		t.Fatalf("Failed to copy kommando binary to example directory: %v", err)
	}

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
