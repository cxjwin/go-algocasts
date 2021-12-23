package leetcode20

// https://leetcode.com/problems/valid-parentheses/description/

func isValid(s string) bool {
	m := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	stack := make([]rune, 0)

	for _, v := range s {
		c, ok := m[v]
		if ok {
			stack = append(stack, c)
		} else {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if v != top {
					return false
				}
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
