package leetcode144

// https://leetcode.com/problems/binary-tree-preorder-traversal/

// TreeNode - Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalRecursive(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append((*res), root.Val)
	preorderTraversalRecursive(root.Left, res)
	preorderTraversalRecursive(root.Right, res)
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	pres := &res
	preorderTraversalRecursive(root, pres)
	return *pres
}

func preorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	res := make([]int, 0)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		// first append, last pop
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		// last append, first pop
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

func preorderTraversalIterative2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	res := make([]int, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		root = node.Right
		stack = stack[:len(stack)-1]
	}

	return res
}
