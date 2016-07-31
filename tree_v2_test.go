package binary_search_tree_test

import (
	"testing"

	"github.com/brettallred/binary_search_tree"
	"github.com/stretchr/testify/assert"
)

func TestTreeV2Insert(t *testing.T) {
	tree := &binary_search_tree.TreeV2{}
	assert := assert.New(t)
	assert.Nil(tree.Root)

	// Root Node
	var value binary_search_tree.IntNode = 8

	tree.Insert(value)
	assert.Equal(tree.Root.Value, value)

	// First Level of the Tree

	value = 3
	tree.Insert(value)
	assert.Equal(tree.Root.Children[0].Value, value)

	value = 10
	tree.Insert(value)
	assert.Equal(tree.Root.Children[1].Value, value)

	// // Second Level of the Tree

	value = 1
	tree.Insert(value)
	assert.Equal(tree.Root.Children[0].Children[0].Value, value)

	value = 6
	tree.Insert(value)
	assert.Equal(tree.Root.Children[0].Children[1].Value, value)

	value = 14
	tree.Insert(value)
	assert.Equal(tree.Root.Children[1].Children[1].Value, value)

	// // Third Level of the Tree
	value = 4
	tree.Insert(value)
	assert.Equal(tree.Root.Children[0].Children[1].Children[0].Value, value)

	value = 7
	tree.Insert(value)
	assert.Equal(tree.Root.Children[0].Children[1].Children[1].Value, value)

	value = 13
	tree.Insert(value)
	assert.Equal(tree.Root.Children[1].Children[1].Children[0].Value, value)
}

func TestTreeV2SearchValuesThatExist(t *testing.T) {
	tree := &binary_search_tree.TreeV2{}
	assert := assert.New(t)
	assert.Nil(tree.Root)

	var value binary_search_tree.IntNode = 8

	tree.Insert(value)
	node := tree.Search(value)
	assert.Equal(node.Value, value)

	value = 3
	tree.Insert(value)
	node = tree.Search(value)
	assert.Equal(node.Value, value)
}

func TestTreeV2SearchValuesThatDoNotExist(t *testing.T) {
	tree := &binary_search_tree.TreeV2{}

	var value binary_search_tree.IntNode = 10
	tree.Insert(value)

	value = 8
	node := tree.Search(value)

	assert := assert.New(t)
	assert.Nil(node)
}
