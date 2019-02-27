package algo

import . "github.com/cxjwin/go-algocasts/datastructure"
import "github.com/cxjwin/go-algocasts/utils"

// https://leetcode.com/problems/balanced-binary-tree/

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getHeight(root.Left)
	right := getHeight(root.Right)
	return utils.IntMax(left, right) + 1
}

func getHeightAndCheck(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getHeightAndCheck(root.Left)
	if left == -1 {
		return -1
	}

	right := getHeightAndCheck(root.Right)
	if right == -1 {
		return -1
	}

	if utils.IntAbs(left-right) > 1 {
		return -1
	}

	return utils.IntMax(left, right) + 1
}

func isBalancedTreeTopDown(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return utils.IntAbs(getHeight(root.Left)-getHeight(root.Right)) <= 1 &&
		isBalancedTreeTopDown(root.Left) &&
		isBalancedTreeTopDown(root.Right)
}

func isBalancedTreeBottomUp(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return getHeightAndCheck(root) != -1
}
