package algo

import (
	"math"

	"github.com/golang-collections/collections/stack"
)

// MinStack min stack struct
type MinStack struct {
	min   *stack.Stack
	stack *stack.Stack
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	s := MinStack{}
	s.min = &stack.Stack{}
	s.stack = &stack.Stack{}
	return s
}

// Push a value
func (s *MinStack) Push(x int) {
	s.stack.Push(x)
	if s.min.Len() == 0 || x < s.min.Peek().(int) {
		s.min.Push(x)
	}
}

// Pop a value
func (s *MinStack) Pop() {
	top := s.stack.Peek()
	if s.min.Len() > 0 && top == s.min.Peek() {
		s.min.Pop()
	}
	s.stack.Pop()
}

// Top value in stack
func (s *MinStack) Top() int {
	top := s.stack.Peek()
	if top == nil {
		return math.MaxInt32
	}
	return top.(int)
}

// GetMin get min value in stack
func (s *MinStack) GetMin() int {
	min := s.min.Peek()
	if min == nil {
		return math.MaxInt32
	}
	return min.(int)
}

// MinStack2 min stack struct use slice
type MinStack2 struct {
	min   int
	stack []int
}

// Constructor2 initialize your data structure here.
func Constructor2() MinStack2 {
	s := MinStack2{}
	s.min = math.MaxInt32
	s.stack = []int{}
	return s
}

// Push push a value
func (s *MinStack2) Push(x int) {
	if x <= s.min {
		// add last min value to slice
		s.stack = append(s.stack, s.min)
		s.min = x
	}
	// add x to slice
	s.stack = append(s.stack, x)
}

// Pop pop a Value
func (s *MinStack2) Pop() {
	l := len(s.stack)
	if l > 0 {
		top := s.stack[l-1]
		if s.min == top && l > 1 {
			// fetch pre min value
			s.min = s.stack[l-2]
			s.stack = s.stack[:l-2]
		} else {
			s.stack = s.stack[:l-1]
		}
	}
}

// Top get the top value
func (s *MinStack2) Top() int {
	l := len(s.stack)
	if l == 0 {
		return math.MaxInt32
	}
	return s.stack[l-1]
}

// GetMin get the min value
func (s *MinStack2) GetMin() int {
	return s.min
}
