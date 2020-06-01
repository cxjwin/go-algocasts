package leetcode236

import . "cxjwin.com/go-algocasts/datastructure"

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else {
		return right
	}
}

func search(root, target *TreeNode, res *[]*TreeNode) bool {
	if root == nil {
		return false
	}

	*res = append(*res, root)
	if root == target {
		return true
	}

	if search(root.Left, target, res) || search(root.Right, target, res) {
		return true
	}

	*res = (*res)[:len(*res)-1]
	return false
}

func lowestCommonAncestorUsingArray(root, p, q *TreeNode) *TreeNode {
	res1 := &[]*TreeNode{}
	res2 := &[]*TreeNode{}
	search(root, p, res1)
	search(root, q, res2)

	n := len(*res1)
	if n > len(*res2) {
		n = len(*res2)
	}

	last := -1
	for i := 0; i < n; i++ {
		if (*res1)[i] == (*res2)[i] {
			last = i
		}
	}
	if last >= 0 {
		return (*res1)[last]
	}

	return nil
}
