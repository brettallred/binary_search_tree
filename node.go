package binary_search_tree

import ()

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

func (node *Node) IsLeaf() bool {
	return node.Children[0] == nil && node.Children[1] == nil
}

func (node *Node) Delete(value int) {

	for i, child := range node.Children {
		if child != nil && child.IntValue == value && child.IsLeaf() {
			node.Children[i] = nil
			break
		}
	}

	var i int

	if node.IntValue > value {
		i = 0
	} else {
		i = 1
	}

	if node.Children[i] != nil {
		node.Children[i].Delete(value)
	}
}
