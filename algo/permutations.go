package algo

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

func permutationsRecursive(nums []int, start int, res [][]int) [][]int {
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		return append(res, temp)
	}

	for i := start; i < len(nums); i++ {
		swap(nums, i, start)
		res = permutationsRecursive(nums, start+1, res)
		swap(nums, i, start)
	}

	return res
}

func permutations(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}

	res := [][]int{}
	res = permutationsRecursive(nums, 0, res)
	return res
}
