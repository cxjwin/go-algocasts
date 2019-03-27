package leetcode448

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

func findDisappearedNumbers(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}

	n := len(nums)
	flag := make([]bool, n)

	for _, v := range nums {
		flag[v-1] = true
	}
	res := make([]int, 0)
	for i, v := range flag {
		if !v {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func findDisappearedNumbersO1(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}

	n := len(nums)
	for i := 0; i < n; i++ {
		idx := abs(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}
