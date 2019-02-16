package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func levelOrder(root *ds.Tree) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)

	queue := []*ds.Tree{root}

	for len(queue) != 0 {
		row := make([]int, 0)
		temp := make([]*ds.Tree, 0)
		for _, v := range queue {
			row = append(row, v.Value.(int))
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		res = append(res, row)
		queue = temp
	}

	return res
}
