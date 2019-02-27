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
