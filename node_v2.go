package binary_search_tree

import ()

type NodeV2 struct {
	Value    NodeValue
	Children [2]*NodeV2
}

type NodeValue interface {
	IsGreaterThan(NodeValue) bool
	Equals(NodeValue) bool
}

type IntNode int

func (node IntNode) IsGreaterThan(value NodeValue) bool {
	return node > value.(IntNode)
}

func (node IntNode) Equals(value NodeValue) bool {
	return node == value.(IntNode)
}

type StringNode int

func (node StringNode) IsGreaterThan(value NodeValue) bool {
	// Put String Comparison Logic Here
	return true
}

func (node StringNode) Equals(value NodeValue) bool {
	// Put String Comparison Logic Here
	return true
}

func (node *NodeV2) Insert(child NodeV2) {
	var i int

	if node.Value.IsGreaterThan(child.Value) {
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

func (node *NodeV2) Search(value NodeValue) *NodeV2 {
	var i int

	if value.Equals(node.Value) {
		return node
	} else if value.IsGreaterThan(node.Value) {
		i = 1
	} else {
		i = 0
	}

	if node.Children[i] == nil {
		return nil
	}

	return node.Children[i].Search(value)
}
