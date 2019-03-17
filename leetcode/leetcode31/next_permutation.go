package leetcode31

// https://leetcode.com/problems/next-permutation/

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func nextPermutation(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	n := len(nums)

	p := n - 2
	for p >= 0 && nums[p] >= nums[p+1] {
		p--
	}

	if p >= 0 {
		i := n - 1
		for i > p && nums[i] <= nums[p] {
			i--
		}
		swap(nums, i, p)
	}

	for i, j := p+1, n-1; i < j; i, j = i+1, j-1 {
		swap(nums, i, j)
	}
}
