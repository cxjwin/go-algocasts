package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func binaryTreeInorderTraversal(root *ds.Tree) []int {
	if root == nil {
		return make([]int, 0)
	}

	left := binaryTreeInorderTraversal(root.Left)
	res := append(left, root.Value.(int))
	right := binaryTreeInorderTraversal(root.Right)
	res = append(res, right...)
	return res
}
