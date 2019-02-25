package algo

// https://leetcode.com/problems/pancake-sorting/

func findMax(nums []int, n int) int {
	mi := 0
	for i := 0; i < n; i++ {
		if nums[i] > nums[mi] {
			mi = i
		}
	}
	return mi
}

func flip(nums []int, idx int) {
	start := 0
	for start < idx {
		temp := nums[start]
		nums[start] = nums[idx]
		nums[idx] = temp
		start++
		idx--
	}
}

func pancakeSort(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}

	n := len(nums)
	res := make([]int, 0)
	for curSize := n; curSize > 1; curSize-- {
		mi := findMax(nums, curSize)

		if mi != curSize-1 {
			res = append(res, mi+1)
			flip(nums, mi)
			res = append(res, curSize)
			flip(nums, curSize-1)
		}
	}

	return res
}
