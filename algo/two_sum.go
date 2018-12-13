package algo

func getTwoNumSumToGivenValueHashMap(nums []int, target int) (int, int) {
	if len(nums) == 0 {
		return -1, -1
	}

	m := make(map[int]int)

	for idx, val := range nums {
		key := target - val
		v, ok := m[key]
		if ok {
			return v, idx
		}
		m[val] = idx
	}

	return -1, -1
}
