package leetcode122

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	n := len(prices)
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, 2)
	}
	d[0][0] = 0
	d[0][1] = -prices[0]

	maxProfit := 0
	for i := 1; i < n; i++ {
		d[i][0] = max(d[i-1][0], d[i-1][1]+prices[i])
		d[i][1] = max(d[i-1][0]-prices[i], d[i-1][1])

		maxProfit = max(maxProfit, d[i][0])
	}

	return maxProfit
}
