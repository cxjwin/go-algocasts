package leetcode450

import . "cxjwin.com/go-algocasts/datastructure"

// https://leetcode.com/problems/delete-node-in-a-bst/

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
	
		leftMax := root.Left
		for leftMax.Right != nil {
			leftMax = leftMax.Right
		}
		leftMax.Right = root.Right
		root = root.Left
	}

	return root
}
