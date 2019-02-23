package algo

import "github.com/cxjwin/go-algocasts/datastructure"

// leetcode : https://leetcode.com/problems/validate-binary-search-tree/

func minTreeNode(root *ds.Tree) *ds.Tree {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func maxTreeNode(root *ds.Tree) *ds.Tree {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func isValidBST(root *ds.Tree) bool {
	if root == nil {
		return true
	}

	leftValid := (root.Left == nil || root.Value.(int) > maxTreeNode(root.Left).Value.(int))
	rightValid := (root.Right == nil || root.Value.(int) < minTreeNode(root.Right).Value.(int))

	return leftValid && rightValid && isValidBST(root.Left) && isValidBST(root.Right)
}

func isValidBSTUseBound(root *ds.Tree) bool {
	return isValidBSTBound(root, nil, nil)
}

func isValidBSTBound(root *ds.Tree, lower *ds.Tree, upper *ds.Tree) bool {
	if root == nil {
		return true
	}
	if lower != nil && lower.Value.(int) > root.Value.(int) {
		return false
	}
	if upper != nil && upper.Value.(int) < root.Value.(int) {
		return false
	}
	return isValidBSTBound(root.Left, lower, root) && isValidBSTBound(root.Right, root, upper)
}
