package algo

// Pair - two value struct
type Pair struct {
	firest, second interface{}
}

func getTwoNumSumToGivenValueHashMap(nums []int, target int) Pair {
	if len(nums) == 0 {
		return Pair{-1, -1}
	}

	m := make(map[int]int)

	for idx, val := range nums {
		key := target - val
		v, ok := m[key]
		if ok {
			return Pair{v, idx}
		}
		m[val] = idx
	}

	return Pair{-1, -1}
}
