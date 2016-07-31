package binary_search_tree

import ()

type Tree struct {
	Root *Node
}

func (tree *Tree) Insert(value int) {
	node := &Node{IntValue: value}

	if tree.Root == nil {
		tree.Root = node
	} else {
		tree.Root.Insert(*node)
	}
}

func (tree *Tree) Search(value int) *Node {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.Search(value)
}

func (tree *Tree) Delete(value int) {
	if tree.Root == nil {
		return
	}

	tree.Root.Delete(value)
}
