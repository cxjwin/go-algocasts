package leetcode746

// https://leetcode.com/problems/min-cost-climbing-stairs/

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	if cost == nil || len(cost) == 0 {
		return 0
	}
	n := len(cost)
	if n == 1 {
		return cost[0]
	}

	d := make([]int, n)
	d[0] = cost[0]
	d[1] = cost[1]
	for i := 2; i < n; i++ {
		d[i] = min(d[i-2], d[i-1]) + cost[i]
	}

	return min(d[n-2], d[n-1])
}

func minCostClimbingStairsO1(cost []int) int {
	if cost == nil || len(cost) == 0 {
		return 0
	}
	n := len(cost)
	if n == 1 {
		return cost[0]
	}

	fir, sec := cost[0], cost[1]
	for i := 2; i < n; i++ {
		cur := min(fir, sec) + cost[i]
		fir = sec
		sec = cur
	}
	return min(fir, sec)
}
