package leetcode112

// https://leetcode.com/problems/path-sum/

// TreeNode - Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) ||
		hasPathSum(root.Right, sum-root.Val)
}

func hasPathSumIterative(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	stack := []*TreeNode{root}
	sumStack := []int{sum}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum := sumStack[len(sumStack)-1]
		sumStack = sumStack[:len(sumStack)-1]
		if node.Left == nil && node.Right == nil && node.Val == sum {
			return true
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
			sumStack = append(sumStack, sum-node.Val)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			sumStack = append(sumStack, sum-node.Val)
		}
	}

	return false
}
