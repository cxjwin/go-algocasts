package leetcode235

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}

	return root
}
