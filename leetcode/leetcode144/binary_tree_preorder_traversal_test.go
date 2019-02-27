package leetcode144

import "testing"

func TestPreorderTraversal(t *testing.T) {
	type testFunc func(root *TreeNode) []int

	testBody := func(f testFunc, t *testing.T) {
		root := &TreeNode{Val: 1, Left: nil, Right: nil}
		root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
		root.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
		root.Left.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
		root.Left.Right.Left = &TreeNode{Val: 5, Left: nil, Right: nil}

		// 1, 2, 4, 5, 3
		res := f(root)
		if res[0] != 1 ||
			res[1] != 2 ||
			res[2] != 4 ||
			res[3] != 5 ||
			res[4] != 3 {
			t.Error("1, 2, 4, 5, 3")
		}
	}

	testBody(preorderTraversal, t)
	testBody(preorderTraversalIterative, t)
	testBody(preorderTraversalIterative2, t)
}
