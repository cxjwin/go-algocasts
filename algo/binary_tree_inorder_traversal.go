package algo

import . "github.com/cxjwin/go-algocasts/datastructure"
import "github.com/golang-collections/collections/stack"

// https://leetcode.com/problems/binary-tree-inorder-traversal/

func binaryTreeInorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	left := binaryTreeInorderTraversal(root.Left)
	res := append(left, root.Val)
	right := binaryTreeInorderTraversal(root.Right)
	res = append(res, right...)
	return res
}

func binaryTreeInorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	stack := new(stack.Stack)
	res := make([]int, 0)

	for root != nil || stack.Len() != 0 {
		for root != nil {
			if root != nil {
				stack.Push(root)
			}
			root = root.Left
		}
		root = stack.Pop().(*TreeNode)
		if root != nil {
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res
}
