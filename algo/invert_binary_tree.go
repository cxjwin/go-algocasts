package algo

import "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/invert-binary-tree/

func invertTreeRecursive(root *ds.Tree) *ds.Tree {
	if root == nil {
		return root
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	invertTreeRecursive(root.Left)
	invertTreeRecursive(root.Right)
	return root
}

func invertTreeInterative(root *ds.Tree) *ds.Tree {
	if root == nil {
		return root
	}

	queue := []*ds.Tree{root}

	for len(queue) != 0 {
		// pop a node
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		// swap
		temp := node.Left
		node.Left = node.Right
		node.Right = temp

		// push
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}
