package leetcode114

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

var prev *TreeNode

func flattenReversePreorder(root *TreeNode) {
	if root == nil {
		return
	}

	flattenReversePreorder(root.Right)
	flattenReversePreorder(root.Left)

	root.Left = nil
	root.Right = prev
	prev = root
}

func flatten(root *TreeNode) {
	prev = nil
	flattenReversePreorder(root)
}

func flattenWithArray(root *TreeNode, arr *[]*TreeNode) {
	if root == nil {
		return
	}

	*arr = append(*arr, root)
	flattenWithArray(root.Left, arr)
	flattenWithArray(root.Right, arr)
}

func flattenPreorder(root *TreeNode) {
	if root == nil {
		return
	}

	temp := make([]*TreeNode, 0)
	arr := &temp
	flattenWithArray(root, arr)

	cur := root
	for i := 1; i < len(*arr); i++ {
		cur.Left = nil
		cur.Right = (*arr)[i]
		cur = cur.Right
	}
}
