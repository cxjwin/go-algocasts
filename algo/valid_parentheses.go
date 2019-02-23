package algo

// leetcode : https://leetcode.com/problems/valid-parentheses/

func isValidParentheses(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else if len(stack) == 0 {
			return false
		} else {
			t := stack[len(stack)-1]
			if v == ')' && t != '(' {
				return false
			}
			if v == ']' && t != '[' {
				return false
			}
			if v == '}' && t != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isValidParenthesesShort(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' {
			stack = append(stack, ')')
		} else if v == '[' {
			stack = append(stack, ']')
		} else if v == '{' {
			stack = append(stack, '}')
		} else if len(stack) == 0 {
			return false
		} else {
			if stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
