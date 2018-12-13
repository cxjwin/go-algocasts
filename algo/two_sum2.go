package algo

func getTwoNumSumToGivenValue(nums []int, target int) (int, int) {
	if len(nums) == 0 {
		return -1, -1
	}

	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			return i + 1, j + 1
		}
	}

	return -1, -1
}
