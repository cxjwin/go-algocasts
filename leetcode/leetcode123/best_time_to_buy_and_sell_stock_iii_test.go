package leetcode123

import "testing"

func TestMaxProfit(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	res := maxProfit(prices)
	if res != 6 {
		t.Error("Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3. Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.")
	}

	prices = []int{1, 2, 3, 4, 5}
	res = maxProfit(prices)
	if res != 4 {
		t.Error("Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4. Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.")
	}

	prices = []int{7, 6, 4, 3, 1}
	res = maxProfit(prices)
	if res != 0 {
		t.Error("In this case, no transaction is done, i.e. max profit = 0.")
	}
}
