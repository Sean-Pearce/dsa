package bst

import "errors"

var (
	ErrorNotFound     = errors.New("not found")
	ErrorAlreadyExist = errors.New("already exist")
)

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

func (b *BSTNode) insert(x int) error {
	if x < b.val {
		if b.left == nil {
			b.left = &BSTNode{val: x}
		} else {
			return b.left.insert(x)
		}
	} else if x > b.val {
		if b.right == nil {
			b.right = &BSTNode{val: x}
		} else {
			return b.right.insert(x)
		}
	} else {
		return ErrorAlreadyExist
	}
	return nil
}

func (b *BSTNode) find(x int) bool {
	if x < b.val && b.left != nil {
		return b.left.find(x)
	} else if x > b.val && b.right != nil {
		return b.right.find(x)
	} else if x == b.val {
		return true
	}
	return false
}

func (b *BSTNode) inorder() []int {
	var res []int
	if b.left != nil {
		res = append(res, b.left.inorder()...)
	}
	res = append(res, b.val)
	if b.right != nil {
		res = append(res, b.right.inorder()...)
	}
	return res
}
