package ds

import "fmt"

// TreeNode - Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTree - new a tree point
func NewTree(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

// Insert - insert node to tree
func (t *TreeNode) Insert(v int, left bool) *TreeNode {
	if t == nil {
		return &TreeNode{v, nil, nil}
	}
	if left {
		t.Left = t.Left.Insert(v, true)
		return t.Left
	}

	t.Right = t.Right.Insert(v, false)
	return t.Right
}

// Desc - description of tree
func (t *TreeNode) Desc() {
	if t == nil {
		return
	}

	fmt.Println(t.Val)
	t.Left.Desc()
	t.Right.Desc()
}
