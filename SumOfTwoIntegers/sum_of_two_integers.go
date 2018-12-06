package main

import "fmt"

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

func main() {
	a, b := 1, 2
	fmt.Printf("recusive - res : %v\n", getSumRecusive(a, b))

	a, b = 9, 11
	fmt.Printf("interative - res : %v\n", getSumInterative(a, b))
}
