package main

import (
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
