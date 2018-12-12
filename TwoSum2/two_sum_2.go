package algo

// Pair - two value struct
type Pair struct {
	first, second interface{}
}

func getTwoNumSumToGivenValue(nums []int, target int) Pair {
	if len(nums) == 0 {
		return Pair{-1, -1}
	}

	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			return Pair{i + 1, j + 1}
		}
	}

	return Pair{-1, -1}
}
