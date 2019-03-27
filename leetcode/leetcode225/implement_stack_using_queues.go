package leetcode225

// https://leetcode.com/problems/implement-stack-using-queues/description/

type queue []int

func (q *queue) push(v int) {
	*q = append(*q, v)
}

func (q *queue) pop() int {
	v := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return v
}

func (q *queue) peek() int {
	return (*q)[len(*q)-1]
}

func (q *queue) size() int {
	return len(*q)
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

type MyStack struct {
	in  queue
	out queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	stack := MyStack{}
	stack.in = make(queue, 0)
	stack.out = make(queue, 0)
	return stack
}

/** Push element x onto stack. */
func (s *MyStack) Push(x int) {
	s.in.push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	for s.in.size() > 1 {
		v := s.in.pop()
		s.out.push(v)
	}

	s.in, s.out = s.out, s.in

	return s.out.pop()
}

/** Get the top element. */
func (s *MyStack) Top() int {
	return s.out.peek()
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return s.in.empty()
}
