func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func _maxProfit(prices []int) int {
	n := len(prices)

	maxProfit := 0

	for i := 1; i < n; i++ {
		temp := prices[i] - prices[i-1]
		if temp > 0 {
			maxProfit += temp
		}
	}

	return maxProfit
}

func maxProfit(t int, prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	n := len(prices)

	if t > n {
		return _maxProfit(prices) // for large t cause runtime error
	}

	K := t + 1

	d := make([][][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([][]int, K)
		for k := 0; k < K; k++ {
			d[i][k] = make([]int, 2)
		}
	}

	for k := 0; k < K; k++ {
		d[0][k][0] = 0
		d[0][k][1] = -prices[0]
	}

	maxProfit := 0

	for i := 1; i < n; i++ {
		for k := 0; k < K; k++ {
			if k > 0 {
				d[i][k][0] = max(
					d[i-1][k][0],             // do nothing
					d[i-1][k-1][1]+prices[i], // sell
				)
			} else {
				d[i][k][0] = d[i-1][k][0]
			}
			d[i][k][1] = max(
				d[i-1][k][1],           // do nothing
				d[i-1][k][0]-prices[i], // buy
			)
		}

		maxProfit = max(maxProfit, d[i][K-1][0])
	}

	return maxProfit
}