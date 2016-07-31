package binary_search_tree

import ()

type TreeV2 struct {
	Root *NodeV2
}

func (tree *TreeV2) Insert(value NodeValue) {
	node := &NodeV2{Value: value}

	if tree.Root == nil {
		tree.Root = node
	} else {
		tree.Root.Insert(*node)
	}
}

func (tree *TreeV2) Search(value NodeValue) *NodeV2 {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.Search(value)
}
