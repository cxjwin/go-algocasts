package leetcode1

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		k := target - num
		v, ok := m[k]
		if ok {
			return []int{v, i}
		}
		m[num] = i
	}
	return []int{-1, -1}
}
