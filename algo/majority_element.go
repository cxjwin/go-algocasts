package algo

// https://leetcode.com/problems/majority-element/

func majorityElementWithMap(nums []int) int {
	m := make(map[int]int)
	max := 1
	res := nums[0]

	for _, v := range nums {
		count, ok := m[v]
		if ok {
			count++
		} else {
			count = 0
		}
		m[v] = count
		if count > max {
			res = v
			max = count
		}
	}

	return res
}

func majorityElement(nums []int) int {
	count := 1
	num := nums[0]
	l := len(nums)

	for i := 1; i < l; i++ {
		if count == 0 {
			num = nums[i]
			count++
		} else if nums[i] == num {
			count++
		} else {
			count--
		}
	}

	return num
}
