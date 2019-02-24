package algo

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfitBruteForce(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

func maxProfitOnePass(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	buy := prices[0]
	max := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else {
			if prices[i]-buy > max {
				max = prices[i] - buy
			}
		}
	}

	return max
}
