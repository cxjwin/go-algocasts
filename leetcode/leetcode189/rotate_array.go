package leetcode189

// https://leetcode.com/problems/rotate-array/

func rotate(nums []int, k int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	n := len(nums)
	m := k % n

	temp := make([]int, n)
	i := 0

	for j := n - m; j < n; j++ {
		temp[i] = nums[j]
		i++
	}
	for j := 0; j < n-m; j++ {
		temp[i] = nums[j]
		i++
	}
	for j := 0; j < n; j++ {
		nums[j] = temp[j]
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func move2Head(nums []int, i int, j int) {
	for j > i {
		swap(nums, j, j-1)
		j--
	}
}

func rotate2(nums []int, k int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	n := len(nums)
	m := k % n

	for i := 0; i < m; i++ {
		move2Head(nums, i, n-m+i)
	}
}

func reverse(nums []int, i int, j int) {
	for ; i < j; i, j = i+1, j-1 {
		swap(nums, i, j)
	}
}

func rotate3(nums []int, k int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	n := len(nums)
	m := k % n

	reverse(nums, 0, n-1)
	reverse(nums, 0, m-1)
	reverse(nums, m, n-1)
}
