# Kommando Repository Instructions for AI Agents

This repository contains the source code for `kommando`, a Go project.

## Project Structure

- `cmd/`: Contains the main application entry point.
- `example/`: Contains example usage of the library.
- `kommando_test.go`: Contains tests for the library.
- `Makefile`: Contains build and test commands.

## Development

### Building and Testing

- **Run Tests**: Use `make test` to build the project and run all tests.
- **Build**: Use `make build` (or `go build`) to compile the project. The binary is output to `build/kommando`.

### Coding Conventions

- Follow standard Go coding conventions (Effective Go).
- Ensure all code is formatted using `gofmt`.
- Ensure `golangci-lint` passes (if available).

## AI Agent Guidelines

- When making changes, always run `make test` to ensure no regressions.
- If adding new features, add corresponding tests.
- Check `go.mod` for dependencies.
