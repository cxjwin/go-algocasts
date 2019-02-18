package algo

import "github.com/cxjwin/go-algocasts/datastructure"
import "github.com/cxjwin/go-algocasts/utils"

func getHeight(root *ds.Tree) int {
	if root == nil {
		return 0
	}

	left := getHeight(root.Left)
	right := getHeight(root.Right)
	return utils.IntMax(left, right) + 1
}

func getHeightAndCheck(root *ds.Tree) int {
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

func isBalancedTreeTopDown(root *ds.Tree) bool {
	if root == nil {
		return true
	}

	return utils.IntAbs(getHeight(root.Left)-getHeight(root.Right)) <= 1 &&
		isBalancedTreeTopDown(root.Left) &&
		isBalancedTreeTopDown(root.Right)
}

func isBalancedTreeBottomUp(root *ds.Tree) bool {
	if root == nil {
		return true
	}

	return getHeightAndCheck(root) != -1
}
