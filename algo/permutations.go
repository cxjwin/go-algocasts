package algo

import "container/list"

func factorial(n int) (result int) {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func permutationsRecursive(nums []int, start int, res *list.List) {
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		res.PushBack(temp)
		return
	}

	for i := start; i < len(nums); i++ {
		swap(nums, i, start)
		permutationsRecursive(nums, start+1, res)
		swap(nums, i, start)
	}
}

func permutations(nums []int) *list.List {
	if nums == nil || len(nums) == 0 {
		return list.New()
	}

	res := list.New()
	permutationsRecursive(nums, 0, res)
	return res
}
