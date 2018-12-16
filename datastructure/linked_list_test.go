package ds

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := List{nil, nil}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)

	head := list.Head
	if head.Value != 1 {
		t.Error("list's head is 1")
	}

	tail := list.Tail
	if tail.Value != 6 {
		t.Error("list's tail is 6")
	}

	if tail.Next != nil {
		t.Error("tail's next is nil")
	}

	list.Desc()
}
