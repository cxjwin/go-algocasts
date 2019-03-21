package leetcode322

// https://algocasts.io/episodes/8eGx43GM

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxVal - max int value
const MaxVal = 1<<31 - 1

func coinChangeDP(coins []int, amount int) int {
	d := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		d[i] = MaxVal
	}

	for i := 1; i <= len(coins); i++ {
		for j := coins[i-1]; j <= amount; j++ {
			if d[j-coins[i-1]] != MaxVal {
				cur := d[j-coins[i-1]] + 1
				d[j] = min(d[j], cur)
			}
		}
	}

	return d[amount]
}
