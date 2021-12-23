package leetcode91

// https://leetcode.com/problems/decode-ways/

func isValidTwoDigit(a byte, b byte) bool {
	return (a == '1' && b <= '9') || (a == '2' && b <= '6')
}

func decode(s string, idx int) int {
	if idx == len(s) {
		return 1
	}
	if idx == len(s)-1 && s[idx] != '0' {
		return 1
	}
	if idx == len(s)-1 && s[idx] == '0' {
		return 0
	}
	if idx > len(s) {
		return 0
	}
	num := 0
	if s[idx] != '0' {
		num += decode(s, idx+1)
	}
	if isValidTwoDigit(s[idx], s[idx+1]) {
		num += decode(s, idx+2)
	}
	return num
}

func numDecoding(s string) int {
	return decode(s, 0)
}

func numDecodingDP(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	d := make([]int, n+1)
	d[0] = 1
	d[1] = 1
	if s[0] == '0' {
		d[1] = 0
	}

	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			d[i] += d[i-1]
		}
		if isValidTwoDigit(s[i-2], s[i-1]) {
			d[i] += d[i-2]
		}
	}

	return d[n]
}

func numDecodingDPO1(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	pre2 := 1
	pre1 := 1
	if s[0] == '0' {
		pre1 = 0
	}
	for i := 2; i <= n; i++ {
		cur := 0
		if s[i-1] != '0' {
			cur += pre1
		}
		if isValidTwoDigit(s[i-2], s[i-1]) {
			cur += pre2
		}
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}
