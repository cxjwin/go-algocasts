package leetcode106

import "testing"

func TestBuildTree(t *testing.T) {
	preorder := []int{2, 1, 8, 4, 16}
	inorder := []int{2, 8, 16, 4, 1}

	root := buildTree(preorder, inorder)

	if root.Val != 1 ||
		root.Left.Val != 2 ||
		root.Right.Val != 4 ||
		root.Right.Left.Val != 8 ||
		root.Right.Right.Val != 16 {
		t.Error("wrong tree build")
	}
}
