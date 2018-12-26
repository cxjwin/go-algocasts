package ds

import "fmt"

// ListNode - node
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// List - list
type List struct {
	Len  int
	Head *ListNode
	Tail *ListNode
}

// Insert - insert a value
func (list *List) Insert(v interface{}) *List {
	if list.Head == nil {
		node := ListNode{Value: v, Next: nil}
		list.Head = &node
		list.Tail = list.Head
		list.Len = 1
		return list
	}

	node := ListNode{Value: v, Next: nil}

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
		fmt.Printf("v: %v\n", node.Value)
		node = node.Next
	}
}
