package algo

func stringToInteger(s string) int {
	if len(s) == 0 {
		return 0
	}

	n := len(s)
	p := 0

	for p < n && s[p] == ' ' {
		p++
	}
	if p == n {
		return 0
	}

	negative := false
	if s[p] == '-' {
		negative = true
		p++
	} else if s[p] == '+' {
		p++
	}
	if p == n {
		return 0
	}

	for p < n && s[p] == '0' {
		p++
	}
	if p == n {
		return 0
	}

	var sum int64
	start := p
	for p < n && '0' <= s[p] && s[p] <= '9' {
		sum *= 10
		sum += int64(s[p] - '0')
		p++
		if p-start > 10 {
			if negative {
				return -1 << 31
			}
			return 1<<31 - 1
		}
	}

	if negative {
		sum = -sum
	}

	if sum < (-1 << 31) {
		return -1 << 31
	}

	if sum > (1<<31 - 1) {
		return 1<<31 - 1
	}

	return int(sum)
}
