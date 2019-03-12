package ds

// Stack - int stack
type Stack []int

// Push - push a int value
func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

// Pop - pop a int value
func (s *Stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

// Top - stack top value
func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

// Empty - check stack is empty
func (s *Stack) Empty() bool {
	return len(*s) == 0
}
