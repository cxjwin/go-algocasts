package leetcode122

import "testing"

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(nums)
	if res != 7 {
		t.Error("Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4. Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.")
	}

	nums = []int{1, 2, 3, 4, 5}
	res = maxProfit(nums)
	if res != 4 {
		t.Error("Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4. Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.")
	}
}
