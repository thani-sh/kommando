package kommando

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

type Node struct {
	Path  string
	IsExe bool
	IsDir bool
	kids  map[string]*Node
}

func newNode(path string) *Node {
	n := &Node{Path: path}
	stats, err := os.Stat(path)
	if err != nil {
		return nil
	}
	n.IsDir = stats.IsDir()
	n.IsExe = !n.IsDir && (stats.Mode()&0111 != 0)
	return n
}

func (n *Node) Exec(args []string, env []string) {
	if n.IsDir {
		n.Help()
		os.Exit(1)
	}
	argv := append([]string{n.Path}, args...)
	err := syscall.Exec(n.Path, argv, env)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func (n *Node) Find(parts []string) *Node {
	if len(parts) == 0 || n.Kids() == nil {
		return nil
	}
	if len(parts) == 1 {
		nodeName := strings.Split(parts[0], ".")[0]
		return n.Kids()[nodeName]
	}
	child := n.Kids()[parts[0]]
	if child == nil {
		return nil
	}
	return child.Find(parts[1:])
}

func (n *Node) Kids() map[string]*Node {
	if n.kids != nil {
		return n.kids
	}
	stats, err := os.Stat(n.Path)
	if err != nil || !stats.IsDir() {
		return nil
	}
	kids := make(map[string]*Node)
	files, err := ioutil.ReadDir(n.Path)
	if err != nil {
		return nil
	}
	for _, file := range files {
		fileName := file.Name()
		filePath := filepath.Join(n.Path, fileName)
		nodeName := strings.Split(fileName, ".")[0]
		kids[nodeName] = newNode(filePath)
	}
	n.kids = kids
	return kids
}

func (n *Node) Help() {
	println("Sub Commands:")
	for name := range n.Kids() {
		println(" * " + name)
	}
}

func (n *Node) Complete(parts []string) []string {
	if n.Kids() == nil {
		return nil
	}

	if len(parts) == 0 {
		var matches []string
		for name, node := range n.Kids() {
			if node.IsDir || node.IsExe {
				matches = append(matches, name)
			}
		}
		return matches
	}

	if len(parts) == 1 {
		prefix := parts[0]
		var matches []string
		for name, node := range n.Kids() {
			if strings.HasPrefix(name, prefix) {
				// Only suggest if it is a directory or executable
				if node.IsDir || node.IsExe {
					matches = append(matches, name)
				}
			}
		}
		return matches
	}

	part := parts[0]
	child := n.Kids()[part]
	if child != nil {
		return child.Complete(parts[1:])
	}

	return nil
}
