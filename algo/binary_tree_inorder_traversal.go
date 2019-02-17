package algo

import "github.com/cxjwin/go-algocasts/datastructure"
import "github.com/golang-collections/collections/stack"

func binaryTreeInorderTraversal(root *ds.Tree) []int {
	if root == nil {
		return make([]int, 0)
	}

	left := binaryTreeInorderTraversal(root.Left)
	res := append(left, root.Value.(int))
	right := binaryTreeInorderTraversal(root.Right)
	res = append(res, right...)
	return res
}

func binaryTreeInorderTraversalIterative(root *ds.Tree) []int {
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
		root = stack.Pop().(*ds.Tree)
		if root != nil {
			res = append(res, root.Value.(int))
			root = root.Right
		}
	}

	return res
}
