package ds

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := List{0, nil, nil}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)

	head := list.Head

	if list.Len != 6 {
		t.Error("list's len is 6")
	}

	if head.Val != 1 {
		t.Error("list's head is 1")
	}

	tail := list.Tail
	if tail.Val != 6 {
		t.Error("list's tail is 6")
	}

	if tail.Next != nil {
		t.Error("tail's next is nil")
	}

	list.Desc()
}

func TestArray2List(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	head := Array2List(nums)

	if head.Val != 1 ||
		head.Next.Val != 2 ||
		head.Next.Next.Val != 3 ||
		head.Next.Next.Next.Val != 4 ||
		head.Next.Next.Next.Next.Val != 5 {
		t.Error("1->2->3->4->5->nil")
	}
}
