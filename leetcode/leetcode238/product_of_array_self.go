package leetcode238

func productExceptSelf(nums []int) []int {
	if nums == nil || len(nums) < 2 {
		return nums
	}

	n := len(nums)

	leftProduct := make([]int, n)
	rightProduct := make([]int, n)
	output := make([]int, n)

	leftProduct[0], rightProduct[n-1] = nums[0], nums[n-1]

	for i := 1; i < n; i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i]
	}
	for j := n - 2; j >= 0; j-- {
		rightProduct[j] = rightProduct[j+1] * nums[j]
	}

	output[0], output[n-1] = rightProduct[1], leftProduct[n-2]
	for i := 1; i < n-1; i++ {
		output[i] = leftProduct[i-1] * rightProduct[i+1]
	}

	return output
}

func productExceptSelfO1(nums []int) []int {
	if nums == nil || len(nums) < 2 {
		return nums
	}

	n := len(nums)
	output := make([]int, n)
	output[n-1] = nums[n-1]

	for j := n - 2; j >= 0; j-- {
		output[j] = output[j+1] * nums[j] // right
	}
	for i := 1; i < n; i++ {
		nums[i] = nums[i-1] * nums[i] // left
	}

	first, last := output[1], nums[n-2]
	for i := 1; i < n-1; i++ {
		output[i] = nums[i-1] * output[i+1]
	}
	output[0], output[n-1] = first, last

	return output
}

func productExceptSelfNotChangeO1(nums []int) []int {
	if nums == nil || len(nums) < 2 {
		return nums
	}

	n := len(nums)
	output := make([]int, n)
	output[0] = 1
	for i := 1; i < n; i++ {
		output[i] = output[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= right
		right *= nums[i]
	}

	return output
}
