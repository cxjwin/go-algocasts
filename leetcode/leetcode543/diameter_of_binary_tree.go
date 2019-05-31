package leetcode543

import (
	. "github.com/cxjwin/go-algocasts/datastructure"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode, d *int) int {
	if root == nil {
		return 0
	}

	left, right := maxDepth(root.Left, d), maxDepth(root.Right, d)
	*d = max(*d, left+right)

	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	n := 0
	maxDepth(root, &n)
	return n
}

func diameterOfBinaryTreeIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diameter := 0
	m := make(map[*TreeNode]int)
	st := []*TreeNode{root}

	for len(st) != 0 {
		node := st[len(st)-1]

		if node.Left != nil {
			if _, ok := m[node.Left]; !ok {
				st = append(st, node.Left)
				continue
			}
		}

		if node.Right != nil {
			if _, ok := m[node.Right]; !ok {
				st = append(st, node.Right)
				continue
			}
		}

		st = st[:len(st)-1]

		left := m[node.Left]
		right := m[node.Right]
		diameter = max(diameter, left+right)
		m[node] = max(left, right) + 1
	}

	return diameter
}
