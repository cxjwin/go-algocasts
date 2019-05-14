package leetcode121

import "testing"

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(nums)
	if res != 5 {
		t.Error("Buy on day 2 (price = 1) and sell on day 5 (price = 6).")
	}

	nums = []int{7, 6, 4, 3, 1}
	res = maxProfit(nums)
	if res != 0 {
		t.Error("In this case, no transaction is done, i.e. max profit = 0.")
	}
}
