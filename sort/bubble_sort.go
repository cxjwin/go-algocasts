package sort

// https://www.geeksforgeeks.org/bubble-sort/

func bubbleSort(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j+1] < nums[j] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
}

func bubbleSort2(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	count := len(nums)
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
}

func cocktailSort(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
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
}
