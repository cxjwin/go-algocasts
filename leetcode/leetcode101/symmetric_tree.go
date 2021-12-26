package leetcode101

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/symmetric-tree/

func isSymmetricLeftAndRight(left *TreeNode, right *TreeNode) bool {
	if left != nil && right != nil {
		return left.Val == right.Val &&
			isSymmetricLeftAndRight(left.Left, right.Right) &&
			isSymmetricLeftAndRight(left.Right, right.Left)
	} else {
		return left == nil && right == nil
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricLeftAndRight(root.Left, root.Right)
}
