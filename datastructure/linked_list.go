package ds

import "fmt"

// Node - node
type Node struct {
	Value interface{}
	Next  *Node
}

// List - list
type List struct {
	Len  int
	Head *Node
	Tail *Node
}

// Insert - insert a value
func (list *List) Insert(v interface{}) *List {
	if list.Head == nil {
		node := Node{Value: v, Next: nil}
		list.Head = &node
		list.Tail = list.Head
		list.Len = 1
		return list
	}

	node := Node{Value: v, Next: nil}

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
