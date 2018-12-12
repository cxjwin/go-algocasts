package algo

func singleNumberWithXOR(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	return sum
}

func singleNumberWithMap(nums []int) int {
	m := make(map[int]int)
	sum := 0
	uniqueSum := 0
	for _, n := range nums {
		_, ok := m[n]
		if !ok {
			sum += n * 2
			m[n] = n
		}
		uniqueSum += n
	}
	return sum - uniqueSum
}
