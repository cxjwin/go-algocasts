package leetcode470

import "math/rand"

func rand7() int {
	return rand.Intn(7) + 1
}

func rand49() int {
	return 7*(rand7()-1) + rand7()
}

func rand10() int {
	x := 41
	for x > 40 {
		x = rand49()
	}
	return x%10 + 1
}
