package leetcode538

import (
	. "github.com/cxjwin/go-algocasts/datastructure"

	"testing"
)

func TestConvertBST(t *testing.T) {
	type testFunc func(*TreeNode) *TreeNode

	testBody := func(f testFunc, t *testing.T) {
		nums := []interface{}{5, 2, 13}
		root := BuildTree(nums)
		root = f(root)
		if root.Val != 18 ||
			root.Left.Val != 20 ||
			root.Right.Val != 13 {
			t.Error("20, 18, 13")
		}
	}

	testBody(convertBST, t)
	testBody(convertBSTIterative, t)
}
