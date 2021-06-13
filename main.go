package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

var (
	ErrMissingCommand = errors.New("unable to find command")
)

func main() {
	argv := os.Args[1:]
	argc := len(argv)
	root, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	for n := argc; n >= 1; n-- {
		commandPath, err := findCommandPath(root, argv[:n])
		if err != nil {
			continue
		}
		if commandPath == os.Args[0] {
			continue
		}
		commandArgs := argv[n:]
		if err := executeCommand(commandPath, commandArgs); err != nil {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
	printHelpMessage(root)
	os.Exit(2)
}

func findCommandPath(root string, segments []string) (string, error) {
	commandPath := filepath.Join(root, filepath.Join(segments...))
	if isFileAvailable(commandPath) {
		return commandPath, nil
	}
	commandDir := filepath.Dir(commandPath)
	if isDirectory(commandDir) {
		commandFile := filepath.Base(commandPath)
		result, err := findCommandInDir(commandDir, commandFile)
		if err == nil {
			return result, nil
		}
	}
	return "", ErrMissingCommand
}

func executeCommand(commandPath string, commandArgs []string) error {
	args := append([]string{commandPath}, commandArgs...)
	return syscall.Exec(commandPath, args, os.Environ())
}

func printHelpMessage(root string) {
	fmt.Println("Unable to find matching command in " + root)
	// TODO: print available commands and help messages
}

func findCommandInDir(dir string, base string) (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		if fileName == base || strings.HasPrefix(fileName, base+".") {
			return filepath.Join(dir, fileName), nil
		}
	}
	return "", ErrMissingCommand
}

func isFileAvailable(path string) bool {
	stats, err := os.Stat(path)
	if err != nil {
		return false
	}
	if stats.IsDir() {
		return false
	}
	return true
}

func isDirectory(path string) bool {
	stats, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stats.IsDir()
}
