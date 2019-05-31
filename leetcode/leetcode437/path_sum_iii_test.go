package leetcode437

import "testing"
import . "github.com/cxjwin/go-algocasts/datastructure"

func TestPathSum(t *testing.T) {
	type testFunc func(*TreeNode, int) int

	testBody := func(f testFunc, t *testing.T) {
		nums := []interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}
		root := BuildTree(nums)
		res := f(root, 8)
		if res != 3 {
			t.Error("1 : 3")
		}
	}

	testBody(pathSum, t)
	testBody(pathSumPrefixSum, t)
}
