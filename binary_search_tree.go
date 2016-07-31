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

type Node struct {
	IntValue int
	Children [2]*Node
}

func (node *Node) Insert(child Node) {
	var i int

	if node.IntValue > child.IntValue {
		i = 0
	} else {
		i = 1
	}

	if node.Children[i] == nil {
		node.Children[i] = &child
	} else {
		node.Children[i].Insert(child)
	}
}

func (node *Node) Search(value int) *Node {
	var i int

	if node.IntValue == value {
		return node
	} else if node.IntValue > value {
		i = 0
	} else {
		i = 1
	}

	if node.Children[i] == nil {
		return nil
	}

	return node.Children[i].Search(value)
}
