package algo

func getSumRecusive(a int, b int) int {
	if b == 0 {
		return a
	}
	return getSumRecusive(a^b, (a&b)<<1)
}

func getSumInterative(a int, b int) int {
	sum := a ^ b
	carry := (a & b) << 1
	for carry != 0 {
		tmp := sum
		sum = sum ^ carry
		carry = (tmp & carry) << 1
	}
	return sum
}
