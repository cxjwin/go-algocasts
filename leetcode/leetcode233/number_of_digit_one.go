package leetcode233

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	for i := 1; i <= n; i++ {
		temp := i
		for temp != 0 {
			if temp%10 == 1 {
				count++
			}
			temp /= 10
		}
	}

	return count
}

func countDigitOneMath(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	factor := 1

	for n/factor != 0 {
		digit := (n / factor) % 10

		high := n / (10 * factor)
		if digit == 0 {
			count += high * factor
		} else if digit == 1 {
			count += high * factor
			count += (n % factor) + 1
		} else {
			count += (high + 1) * factor
		}

		factor *= 10
	}

	return count
}
