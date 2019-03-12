package leetcode946

// https://leetcode.com/problems/validate-stack-sequences/

func validateStackSequences(pushed []int, popped []int) bool {
	if pushed == nil || popped == nil || len(pushed) != len(popped) {
		return false
	}

	n := len(pushed)

	stack := make([]int, n)
	top := -1
	j := 0

	for _, v := range pushed {
		top++
		stack[top] = v
		for top != -1 && stack[top] == popped[j] {
			j++
			top--
		}
	}

	return top == -1
}
