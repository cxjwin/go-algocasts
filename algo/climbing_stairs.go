package algo

func climbStairsRecursive(n int) int {
	if n < 2 {
		return 1
	}
	return climbStairsRecursive(n-1) + climbStairsRecursive(n-2)
}

func climbStairsIterative(n int) int {
	first, secend := 1, 1

	for i := 1; i < n; i++ {
		third := first + secend
		first = secend
		secend = third
	}

	return secend
}
