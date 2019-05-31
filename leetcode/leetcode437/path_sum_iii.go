package leetcode437

import . "github.com/cxjwin/go-algocasts/datastructure"

func pathFrom(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	cnt := 0
	if root.Val == sum {
		cnt++
	}
	cnt += pathFrom(root.Left, sum-root.Val)
	cnt += pathFrom(root.Right, sum-root.Val)
	return cnt
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathFrom(root, sum) +
		pathSum(root.Left, sum) +
		pathSum(root.Right, sum)
}

func dfs(root *TreeNode, cur int, sum int, m map[int]int) int {
	if root == nil {
		return 0
	}

	cur += root.Val

	cnt, ok := m[cur-sum]
	if !ok {
		m[cur-sum] = 0
		cnt = 0
	}

	_, ok = m[cur]
	if ok {
		m[cur]++
	} else {
		m[cur] = 1
	}

	cnt += dfs(root.Left, cur, sum, m)
	cnt += dfs(root.Right, cur, sum, m)

	m[cur]--

	return cnt
}

func pathSumPrefixSum(root *TreeNode, sum int) int {
	m := make(map[int]int)
	m[0] = 1
	return dfs(root, 0, sum, m)
}
