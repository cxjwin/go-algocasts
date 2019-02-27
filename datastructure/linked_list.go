package ds

import "fmt"

// ListNode - node
type ListNode struct {
	Val  int
	Next *ListNode
}

// List - list
type List struct {
	Len  int
	Head *ListNode
	Tail *ListNode
}

// Insert - insert a value
func (list *List) Insert(v int) *List {
	if list.Head == nil {
		node := ListNode{Val: v, Next: nil}
		list.Head = &node
		list.Tail = list.Head
		list.Len = 1
		return list
	}

	node := ListNode{Val: v, Next: nil}

	tail := list.Tail
	tail.Next = &node
	list.Tail = tail.Next
	list.Len++
	return list
}

// Desc - description of list
func (list *List) Desc() {
	if list == nil || list.Head == nil {
		fmt.Printf("list is nil")
	}

	node := list.Head
	for node != nil {
		fmt.Printf("v: %v\n", node.Val)
		node = node.Next
	}
}
