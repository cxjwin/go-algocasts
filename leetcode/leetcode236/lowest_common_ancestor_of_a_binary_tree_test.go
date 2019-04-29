package leetcode236

import "testing"
import . "github.com/cxjwin/go-algocasts/datastructure"

func TestLowestCommonAncestor(t *testing.T) {
	type testFunc func(root, p, q *TreeNode) *TreeNode

	testBody := func(f testFunc, t *testing.T) {
		arr := []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
		root := BuildTree(arr)
		node5 := root.Left
		node1 := root.Right
		res := f(root, node5, node1)
		if res.Val != 3 {
			t.Error("The LCA of nodes 5 and 1 is 3.")
		}
		node4 := root.Left.Right.Right
		res = f(root, node5, node4)
		if res.Val != 5 {
			t.Error("The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.")
		}
	}

	// testBody(lowestCommonAncestor, t)
	testBody(lowestCommonAncestorUsingArray, t)
}
