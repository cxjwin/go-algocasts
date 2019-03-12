package leetcode232

// https://leetcode.com/problems/implement-queue-using-stacks/

type stack []int

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

// MyQueue - struct
type MyQueue struct {
	in  stack
	out stack
}

// Constructor - Initialize your data structure here.
func Constructor() MyQueue {
	in := make(stack, 0)
	out := make(stack, 0)
	return MyQueue{in: in, out: out}
}

func (q *MyQueue) checkOutStack() {
	if q.out.empty() {
		for !q.in.empty() {
			x := q.in.pop()
			q.out.push(x)
		}
	}
}

// Push element x to the back of queue.
func (q *MyQueue) Push(x int) {
	q.in.push(x)
}

// Pop - Removes the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	q.checkOutStack()
	return q.out.pop()
}

// Peek - Get the front element.
func (q *MyQueue) Peek() int {
	q.checkOutStack()
	return q.out.top()
}

// Empty - Returns whether the queue is empty.
func (q *MyQueue) Empty() bool {
	return q.in.empty() && q.out.empty()
}
