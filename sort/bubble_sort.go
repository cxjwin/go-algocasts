package sort

func bubbleSort(nums []int) []int {
	if nums == nil {
		return nums
	}

	count := len(nums)
	if count == 0 {
		return nums
	}

	for i := 0; i < count-1; i++ {
		for j := 0; j < count-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}

	return nums
}

func bubbleSort2(nums []int) []int {
	if nums == nil {
		return nums
	}

	count := len(nums)
	if count == 0 {
		return nums
	}

	change := true
	for i := 0; i < count-1 && change; i++ {
		change = false
		for j := 0; j < count-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
				change = true
			}
		}
	}

	return nums
}

func cocktailSort(nums []int) []int {
	if nums == nil {
		return nums
	}

	if len(nums) == 0 {
		return nums
	}

	count := len(nums)
	left := 0
	right := count - 1
	for left < right {
		for j := left; j < right; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
		left++
		for j := right; j > left; j-- {
			if nums[j-1] > nums[j] {
				temp := nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = temp
			}
		}
		right--
	}

	return nums
}
