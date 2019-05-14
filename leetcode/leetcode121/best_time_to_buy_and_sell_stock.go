package leetcode121

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

	maxProfit := 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-buy)
		}
	}
	return maxProfit
}
