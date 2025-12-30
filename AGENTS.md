# Kommando Repository Instructions for AI Agents

This repository contains the source code for `kommando`, a Go project.

## Project Structure

- `cmd/`: Contains the main application entry point.
- `example/`: Contains example usage of the library.
- `kommando_test.go`: Contains tests for the library.
## Development

### Building and Testing

- **Run Tests**: Use `go test ./...` to run all tests. The tests will automatically build the binary required for integration testing.
- **Build**: Use `go build -o build/kommando ./cmd/kommando/...` to compile the project.

### Coding Conventions

- Follow standard Go coding conventions (Effective Go).
- Ensure all code is formatted using `gofmt`.
- Ensure `golangci-lint` passes (if available).

## AI Agent Guidelines

- When making changes, always run `go test ./...` to ensure no regressions.
- If adding new features, add corresponding tests.
- Check `go.mod` for dependencies.
