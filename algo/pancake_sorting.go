package algo

import "fmt"

func find(nums []int, target int) (int, int) {
	idx := -1
	min := target
	for _, v := range nums {
		if v > target {
			min = v
			break
		}
	}

	for i, v := range nums {
		if v < min && v > target {
			idx = i
			min = v
		}
	}

	return idx, min
}

func flip(nums []int, idx int) {
	mid := idx / 2
	for i := 0; i < mid; i++ {
		j := idx - i
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}
}

func pancakeSort(A []int) []int {
	if A == nil || len(A) == 0 {
		return []int{}
	}

	res := make([]int, len(A))
	min := 0
	for i := 0; i < len(A); i++ {
		idx, val := find(A, min)
		fmt.Println(val)
		min = val
		flip(A, idx)
		fmt.Println(A)
		res[i] = idx
	}

	return res
}
