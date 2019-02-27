package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

// leetcode : https://leetcode.com/problems/validate-binary-search-tree/

func minTreeNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func maxTreeNode(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftValid := (root.Left == nil || root.Val > maxTreeNode(root.Left).Val)
	rightValid := (root.Right == nil || root.Val < minTreeNode(root.Right).Val)

	return leftValid && rightValid && isValidBST(root.Left) && isValidBST(root.Right)
}

func isValidBSTUseBound(root *TreeNode) bool {
	return isValidBSTBound(root, nil, nil)
}

func isValidBSTBound(root *TreeNode, lower *TreeNode, upper *TreeNode) bool {
	if root == nil {
		return true
	}
	if lower != nil && lower.Val > root.Val {
		return false
	}
	if upper != nil && upper.Val < root.Val {
		return false
	}
	return isValidBSTBound(root.Left, lower, root) && isValidBSTBound(root.Right, root, upper)
}
