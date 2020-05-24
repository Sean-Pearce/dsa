package bst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBST(t *testing.T) {
	root := &BSTNode{val: 2}

	require.Nil(t, root.insert(1))
	require.Nil(t, root.insert(3))

	require.False(t, root.find(0))
	require.True(t, root.find(1))

	require.Equal(t, []int{1, 2, 3}, root.inorder())
}
