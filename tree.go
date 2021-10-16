package kommando

type Tree struct {
	*Node
}

func NewTree(path string) *Tree {
	return &Tree{Node: newNode(path)}
}
