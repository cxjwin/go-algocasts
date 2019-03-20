package leetcode518

// https://leetcode.com/problems/coin-change-2/

func _change(amount int, start int, coins []int) int {
	if amount == 0 {
		return 1
	}
	if amount < 0 {
		return 0
	}

	res := 0
	for i := start; i < len(coins); i++ {
		res += _change(amount-coins[i], i, coins)
	}
	return res
}

func change(amount int, coins []int) int {
	return _change(amount, 0, coins)
}

func changeDP(amount int, coins []int) int {
	d := make([]int, amount+1)
	d[0] = 1

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			cur := 0
			if j >= coins[i-1] {
				cur = d[j-coins[i-1]]
			}
			d[j] += cur
		}
	}

	return d[amount]
}
