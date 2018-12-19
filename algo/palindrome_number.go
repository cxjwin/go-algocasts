package algo

import "strconv"

func isPalindromeNumberUseString(num int) bool {
	str := strconv.Itoa(num)
	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] == str[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func isPalindromeNumber(num int) bool {
	if num < 0 {
		return false
	}

	temp := num
	// TODO: convert int to long
	res := 0
	for temp != 0 {
		mod := temp % 10
		res = res*10 + mod
		temp /= 10
	}
	return res == num
}
