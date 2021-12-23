package leetcode113

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/path-sum-ii/

func path(root *TreeNode, sum int, elems *[]int, res *[][]int) {
	if root == nil {
		return
	}

	*elems = append(*elems, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		temp := make([]int, len(*elems))
		copy(temp, *elems)
		*res = append(*res, temp)
	}

	path(root.Left, sum-root.Val, elems, res)
	path(root.Right, sum-root.Val, elems, res)
	*elems = (*elems)[:len(*elems)-1]
}

func pathSum(root *TreeNode, sum int) [][]int {
	temp1 := make([]int, 0)
	elems := &temp1
	temp2 := make([][]int, 0)
	res := &temp2
	path(root, sum, elems, res)
	return *res
}

func pathSumIterativePrev(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	curSum := 0
	res := make([][]int, 0)
	elems := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var prev *TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			curSum += root.Val
			stack = append(stack, root)
			elems = append(elems, root.Val)
			root = root.Left
		}

		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == prev {
			if node.Left == nil && node.Right == nil && curSum == sum {
				temp := make([]int, len(elems))
				copy(temp, elems)
				res = append(res, temp)
			}
			curSum -= node.Val
			stack = stack[:len(stack)-1]
			elems = elems[:len(elems)-1]
			prev = node
			root = nil
		} else {
			root = node.Right
		}
	}

	return res
}
