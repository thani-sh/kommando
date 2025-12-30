package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/thani-sh/kommando"
)

func main() {
	argv := os.Args[1:]
	argc := len(argv)
	root, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	tree := kommando.NewTree(root)

	if argc > 0 {
		if argv[0] == "--kommando-completion" {
			completions := tree.Complete(argv[1:])
			for _, c := range completions {
				fmt.Println(c)
			}
			os.Exit(0)
		}

		if argv[0] == "--kommando-init-bash" {
			name := "kommando"
			if argc > 1 {
				name = argv[1]
			}
			printBashInit(name)
			os.Exit(0)
		}

		if argv[0] == "--kommando-init-zsh" {
			name := "kommando"
			if argc > 1 {
				name = argv[1]
			}
			printZshInit(name)
			os.Exit(0)
		}
	}

	for n := argc; n >= 1; n-- {
		path := argv[:n]
		node := tree.Find(path)
		if node == nil || node.Path == os.Args[0] {
			continue
		}
		node.Exec(argv[n:], os.Environ())
	}
	tree.Help()
	os.Exit(1)
}

func printBashInit(name string) {
	script := `
_%[1]s_completion() {
    local cur prev words cword
    _init_completion -n : || return

    local cmd="${words[0]}"
    local args=("${words[@]:1}")

    COMPREPLY=( $( "$cmd" --kommando-completion "${args[@]}" ) )
}
complete -F _%[1]s_completion %[1]s
`
	fmt.Printf(script, name)
}

func printZshInit(name string) {
	script := `
#compdef %[1]s

_%[1]s() {
    local -a completions
    local -a cmd_args
    cmd_args=("${words[@]}")

    # Call the command with special flag to get completions
    # We pass the arguments starting from index 1 (skip command name)
    # in zsh $words includes command name at index 1

    # We need to pass all args to --kommando-completion
    # Zsh arrays are 1-based. words[1] is the command.
    # We want words[2] to the end.

    completions=($($words[1] --kommando-completion "${words[2,-1]}"))

    compadd -a completions
}

compdef _%[1]s %[1]s
`
	fmt.Printf(script, name)
}
