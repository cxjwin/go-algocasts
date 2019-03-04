package leetcode617

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/merge-two-binary-trees/

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 != nil && t2 == nil {
		return t1
	}

	if t1 == nil && t2 != nil {
		return t2
	}

	node := &TreeNode{Val: (t1.Val + t2.Val)}
	node.Left = mergeTrees(t1.Left, t2.Left)
	node.Right = mergeTrees(t1.Right, t2.Right)
	return node
}
