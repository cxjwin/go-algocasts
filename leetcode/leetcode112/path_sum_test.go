package leetcode112

import "testing"

func TestPathSum(t *testing.T) {
	root := &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 11, Left: nil, Right: nil}
	root.Left.Left.Left = &TreeNode{Val: 7, Left: nil, Right: nil}
	root.Left.Left.Right = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 8, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 13, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right.Right.Right = &TreeNode{Val: 1, Left: nil, Right: nil}

	res := hasPathSum(root, 22)
	if !res {
		t.Error("1 - has path sum 22")
	}

	res = hasPathSumIterative(root, 22)
	if !res {
		t.Error("2 - has path sum 22")
	}
}
