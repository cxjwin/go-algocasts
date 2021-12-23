package leetcode199

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/binary-tree-right-side-view/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	queue := []*TreeNode{root}

	for len(queue) != 0 {

		// add this row last val to results
		lastNode := queue[len(queue)-1]
		res = append(res, lastNode.Val)

		// inqueue next row
		temp := make([]*TreeNode, 0)
		for _, v := range queue {
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		queue = temp
	}

	return res
}
