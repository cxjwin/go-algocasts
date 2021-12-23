package leetcode538

import (
	. "github.com/cxjwin/go-algocasts/datastructure"
)

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}

	cur := dfs(root.Right, sum)
	root.Val += cur
	return dfs(root.Left, root.Val)
}

func convertBST(root *TreeNode) *TreeNode {
	dfs(root, 0)
	return root
}

func convertBSTIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	cur := root
	st := make([]*TreeNode, 0)
	sum := 0

	for cur != nil || len(st) != 0 {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Right
		}

		cur = st[len(st)-1]
		st = st[:len(st)-1]

		cur.Val += sum
		sum = cur.Val

		cur = cur.Left
	}

	return root
}
