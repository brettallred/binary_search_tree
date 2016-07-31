package binary_search_tree_test

import (
	"testing"

	"github.com/brettallred/binary_search_tree"
	"github.com/stretchr/testify/assert"
)

func TestIsLeaf(t *testing.T) {
	node := &binary_search_tree.Node{}

	assert.True(t, node.IsLeaf())

	node.Children[0] = node
	assert.False(t, node.IsLeaf())
}
