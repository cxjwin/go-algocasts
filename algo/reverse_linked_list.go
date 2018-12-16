package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func reverseListNodeRecursive(node *ds.Node) *ds.Node {
	if node == nil || node.Next == nil {
		return node
	}

	n := reverseListNodeRecursive(node.Next)
	node.Next.Next = node
	node.Next = nil

	return n
}

func reverseListRecursive(list *ds.List) *ds.List {
	if list == nil || list.Head == nil {
		return list
	}

	list.Tail = list.Head
	list.Head = reverseListNodeRecursive(list.Head)

	return list
}

func reverseListIterative(list *ds.List) *ds.List {
	if list == nil || list.Head == nil {
		return list
	}

	list.Tail = list.Head

	cur := list.Head
	var pre *ds.Node

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	list.Head = pre

	return list
}
