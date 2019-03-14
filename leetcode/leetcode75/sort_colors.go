package leetcode75

// https://leetcode.com/problems/sort-colors/

func sortColors(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	sum0, sum1 := 0, 0
	for _, v := range nums {
		if v == 0 {
			sum0++
		} else if v == 1 {
			sum1++
		}
	}

	n := len(nums)
	for i := 0; i < n; i++ {
		if i < sum0 {
			nums[i] = 0
		} else if i < (sum0 + sum1) {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func sortColors2(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	n := len(nums)
	start, mid, end := 0, 0, n-1
	for mid <= end {
		if nums[mid] == 0 {
			swap(nums, start, mid)
			start++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			swap(nums, mid, end)
			end--
		}
	}
}
