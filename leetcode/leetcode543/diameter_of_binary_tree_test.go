package leetcode543

import "testing"
import . "github.com/cxjwin/go-algocasts/datastructure"

func TestDiameterOfBinaryTree(t *testing.T) {
	type testFunc func(*TreeNode) int

	testBody := func(f testFunc, t *testing.T) {
		nums := []interface{}{1, 2, 3, 4, 5}
		root := BuildTree(nums)
		res := f(root)
		if res != 3 {
			t.Error("Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].")
		}
	}

	// testBody(diameterOfBinaryTree, t)
	testBody(diameterOfBinaryTreeIterative, t)
}
