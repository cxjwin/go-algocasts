package sort

// https://www.geeksforgeeks.org/insertion-sort/

func insertionSort(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	n := len(nums)
	for i := 1; i < n; i++ {
		num := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if num < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = num
	}
}
