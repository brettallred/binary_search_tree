package binary_search_tree_test

import (
	"testing"

	"github.com/brettallred/binary_search_tree"
	"github.com/stretchr/testify/assert"
)

func TestTreeInsert(t *testing.T) {
	tree := &binary_search_tree.Tree{}
	assert := assert.New(t)
	assert.Nil(tree.Root)

	// Root Node
	tree.Insert(8)
	assert.Equal(tree.Root.IntValue, 8)

	// First Level of the Tree
	tree.Insert(3)
	assert.Equal(tree.Root.Children[0].IntValue, 3)

	tree.Insert(10)
	assert.Equal(tree.Root.Children[1].IntValue, 10)

	// Second Level of the Tree
	tree.Insert(1)
	assert.Equal(tree.Root.Children[0].Children[0].IntValue, 1)

	tree.Insert(6)
	assert.Equal(tree.Root.Children[0].Children[1].IntValue, 6)

	tree.Insert(14)
	assert.Equal(tree.Root.Children[1].Children[1].IntValue, 14)

	// Third Level of the Tree
	tree.Insert(4)
	assert.Equal(tree.Root.Children[0].Children[1].Children[0].IntValue, 4)

	tree.Insert(7)
	assert.Equal(tree.Root.Children[0].Children[1].Children[1].IntValue, 7)

	tree.Insert(13)
	assert.Equal(tree.Root.Children[1].Children[1].Children[0].IntValue, 13)

}
