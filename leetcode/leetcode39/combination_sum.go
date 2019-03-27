package leetcode39

// https://leetcode.com/problems/combination-sum/

func sum(candidates []int, target int, start int, elem *[]int, result *[][]int) {
	if target == 0 {
		temp := make([]int, len(*elem))
		copy(temp, *elem)
		*result = append(*result, temp)
		return
	}
	if target < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		*elem = append(*elem, candidates[i])
		sum(candidates, target-candidates[i], i, elem, result)
		*elem = (*elem)[:len(*elem)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	elem := &[]int{}
	result := &[][]int{}
	sum(candidates, target, 0, elem, result)
	return (*result)
}
