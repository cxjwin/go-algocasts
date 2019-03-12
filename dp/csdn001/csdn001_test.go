package csdn001

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	type testFunc func(n int) int

	testBody := func(f testFunc, t *testing.T) {
		res := f(6)
		if res != 8 {
			t.Error(f, "fibonacci 6 is 8")
		}
	}

	testBody(fibonacci, t)
	testBody(fibonacciMemo, t)
	testBody(fibonacciDP, t)
}

func TestCut(t *testing.T) {
	p := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	res := make([]int, len(p)-1)
	for i := 1; i <= 10; i++ {
		res[i-1] = cut(p, i)
	}
	fmt.Println(res)
	for i := 1; i <= 10; i++ {
		res[i-1] = cutDP(p, i)
	}
	fmt.Println(res)
	for i := 1; i <= 10; i++ {
		res[i-1] = cutMemo(p, i)
	}
	fmt.Println(res)
}
